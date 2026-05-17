#!/usr/bin/env python3
"""Convert qiuxin_aliyun_back knowledge base to Self-improvement Q&A format.

Uses Claude API to read .md files and extract Q&A pairs in #q/#a format.
Processes files one at a time with resume-from-checkpoint support.
"""

import os
import json
import sys
import time
from datetime import datetime

# Load .env if available
try:
    from dotenv import load_dotenv
    load_dotenv()
except ImportError:
    pass

from anthropic import Anthropic, APIStatusError, RateLimitError

from prompts import get_prompt, extract_domain, SYSTEM_PROMPT

# ── Configuration ──────────────────────────────────────────────

SOURCE_ROOT = os.path.join(os.path.dirname(os.path.abspath(__file__)), '..', 'qiuxin_aliyun_back')
OUTPUT_ROOT = os.path.join(os.path.dirname(os.path.abspath(__file__)), 'questions')
PROGRESS_FILE = os.path.join(os.path.dirname(os.path.abspath(__file__)), '.progress.json')
MAX_FILE_CHARS = 8000         # truncate files larger than this
MIN_FILE_CHARS = 50           # skip files shorter than this
SLEEP_ON_RATE_LIMIT = 30      # seconds to wait on 429
MAX_RETRIES = 3

EXCLUDE_NAMES = {'投资.md', '服务器列表.md', 'image.png', '.gitkeep'}
EXCLUDE_PATHS = {'99_archive', '.git', '.claude', '.ralph', '.vscode', '.devcontainer'}

MODEL = os.environ.get('ANTHROPIC_MODEL', 'claude-sonnet-4-20250514')

_api_key = os.environ.get('ANTHROPIC_AUTH_TOKEN') or os.environ.get('ANTHROPIC_API_KEY')
_base_url = os.environ.get('ANTHROPIC_BASE_URL')

client_kwargs = {'api_key': _api_key}
if _base_url:
    client_kwargs['base_url'] = _base_url
client = Anthropic(**client_kwargs)

# ── Progress tracking ──────────────────────────────────────────

def load_progress():
    if os.path.exists(PROGRESS_FILE):
        with open(PROGRESS_FILE) as f:
            return json.load(f)
    return {"completed": {}, "total_qa": 0, "started_at": datetime.now().isoformat()}

def save_progress(progress):
    progress["updated_at"] = datetime.now().isoformat()
    with open(PROGRESS_FILE, 'w') as f:
        json.dump(progress, f, indent=2, ensure_ascii=False)

# ── File collection ────────────────────────────────────────────

def collect_files():
    """Walk SOURCE_ROOT and return list of (rel_path, abs_path) to process."""
    files = []
    for dirpath, dirnames, filenames in os.walk(SOURCE_ROOT):
        # Skip excluded directories
        dirnames[:] = [d for d in dirnames if not any(
            ex in os.path.relpath(os.path.join(dirpath, d), SOURCE_ROOT).split(os.sep)
            for ex in EXCLUDE_PATHS
        )]
        for fname in filenames:
            if fname in EXCLUDE_NAMES:
                continue
            if not fname.endswith('.md'):
                continue
            abs_path = os.path.join(dirpath, fname)
            rel_path = os.path.relpath(abs_path, SOURCE_ROOT)
            files.append((rel_path, abs_path))
    return files

def read_file_content(filepath):
    """Read file, return (content, size_chars). Returns (None, 0) on error."""
    try:
        with open(filepath, 'r', encoding='utf-8', errors='replace') as f:
            content = f.read()
        content = content.strip()
        if len(content) < MIN_FILE_CHARS:
            return None, len(content)
        if len(content) > MAX_FILE_CHARS:
            content = content[:MAX_FILE_CHARS] + '\n\n[内容过长，已截断...]'
        return content, len(content)
    except Exception:
        return None, 0

# ── API calling ─────────────────────────────────────────────────

def call_claude(system_prompt, user_prompt):
    """Call Claude API with retry logic. Returns response text or None."""
    for attempt in range(MAX_RETRIES):
        try:
            response = client.messages.create(
                model=MODEL,
                max_tokens=4096,
                system=[
                    {"type": "text", "text": system_prompt,
                     "cache_control": {"type": "ephemeral"}}
                ],
                messages=[{"role": "user", "content": user_prompt}],
            )
            # Extract text from response, handling models that return ThinkingBlock
            for block in response.content:
                if hasattr(block, 'text') and block.text:
                    return block.text
            # Fallback: try the first content block
            if response.content:
                return str(response.content[0])
            return None
        except RateLimitError:
            wait = SLEEP_ON_RATE_LIMIT * (attempt + 1)
            print(f"  Rate limited, waiting {wait}s...")
            if attempt < MAX_RETRIES - 1:
                time.sleep(wait)
        except APIStatusError as e:
            print(f"  API error ({e.status_code})")
            if attempt < MAX_RETRIES - 1:
                time.sleep(5 * (attempt + 1))
        except Exception as e:
            print(f"  Unexpected error: {e}")
            if attempt < MAX_RETRIES - 1:
                time.sleep(5)
    return None

# ── Q&A parsing ─────────────────────────────────────────────────

def parse_qa_from_response(response_text):
    """Parse Claude's response into list of (question, answer) tuples.

    Handles two formats:
    1. Multiline:  # q\\n<text>\\n# a\\n<text>
    2. Inline:     # q <text>\\n# a <text>
    """
    lines = response_text.split('\n')
    q_positions = []
    a_positions = []

    for i, line in enumerate(lines):
        stripped = line.strip()
        if stripped == '# q' or stripped.startswith('# q '):
            q_positions.append(i)
        elif stripped == '# a' or stripped.startswith('# a '):
            a_positions.append(i)

    pairs = []
    for i, q_pos in enumerate(q_positions):
        if i >= len(a_positions):
            break
        a_pos = a_positions[i]

        q_line = lines[q_pos].strip()
        a_line = lines[a_pos].strip()

        # Extract question text
        if q_line == '# q':
            # Multiline: text is between # q line and # a line
            q_text = '\n'.join(lines[q_pos+1:a_pos]).strip()
        else:
            # Inline: text follows '# q ' on same line
            q_text = q_line[4:].strip()

        # Extract answer text
        next_q = q_positions[i+1] if i+1 < len(q_positions) else len(lines)
        if a_line == '# a':
            # Multiline: text is between # a line and next # q
            a_text = '\n'.join(lines[a_pos+1:next_q]).strip()
        else:
            # Inline: text follows '# a ' on same line
            a_text = a_line[4:].strip()
            extra_lines = '\n'.join(lines[a_pos+1:next_q]).strip()
            if extra_lines:
                a_text += '\n' + extra_lines

        if q_text and a_text:
            pairs.append((q_text, a_text))

    return pairs

# ── Output ──────────────────────────────────────────────────────

def get_output_path(rel_path, domain):
    """Map source path to output path under questions/<domain>/."""
    safe_name = rel_path.replace('/', '_').replace(' ', '_')
    domain_dir = os.path.join(OUTPUT_ROOT, domain)
    os.makedirs(domain_dir, exist_ok=True)
    return os.path.join(domain_dir, safe_name)

def write_qa_file(filepath, pairs):
    """Write Q&A pairs in # q / # a format."""
    with open(filepath, 'w', encoding='utf-8') as f:
        for question, answer in pairs:
            f.write(f'# q\n{question}\n# a\n{answer}\n\n')

# ── Single file processing ─────────────────────────────────────

def process_file(rel_path, abs_path, progress):
    """Process a single file: read → API call → parse → write output."""
    domain = extract_domain(rel_path)

    content, size = read_file_content(abs_path)
    if content is None:
        status = "too_short" if size < MIN_FILE_CHARS else "error"
        progress["completed"][rel_path] = {"status": status, "qa_count": 0, "domain": domain}
        return 0

    print(f"  [{domain}] {rel_path} ({len(content)} chars)...", end=' ', flush=True)

    system_prompt, user_prompt = get_prompt(rel_path, domain, content)
    response_text = call_claude(system_prompt, user_prompt)

    if response_text is None:
        progress["completed"][rel_path] = {"status": "api_error", "qa_count": 0, "domain": domain}
        print("FAIL")
        return 0

    pairs = parse_qa_from_response(response_text)

    if pairs:
        output_path = get_output_path(rel_path, domain)
        write_qa_file(output_path, pairs)
        progress["completed"][rel_path] = {
            "status": "done",
            "qa_count": len(pairs),
            "output": str(output_path),
            "domain": domain,
        }
        print(f"{len(pairs)} Q&A")
        return len(pairs)
    else:
        progress["completed"][rel_path] = {
            "status": "no_qa",
            "qa_count": 0,
            "domain": domain,
        }
        print("no Q&A")
        return 0

# ── Main ────────────────────────────────────────────────────────

def main():
    print("=== KB-to-SR Converter ===\n")

    if not _api_key:
        print("ERROR: ANTHROPIC_AUTH_TOKEN or ANTHROPIC_API_KEY not set.")
        sys.exit(1)

    all_files = collect_files()
    print(f"Found {len(all_files)} .md files in {SOURCE_ROOT}\n")

    progress = load_progress()
    completed = set(progress["completed"].keys())
    pending = [(rel, abs) for rel, abs in all_files if rel not in completed]

    print(f"Already processed: {len(completed)}")
    print(f"Remaining: {len(pending)}\n")

    if not pending:
        print("All files processed. Done!")
        return

    total_qa = progress.get("total_qa", 0)
    processed = 0
    errors = 0

    for i, (rel_path, abs_path) in enumerate(pending):
        print(f"[{i+1}/{len(pending)}]", end=' ')

        qa_count = process_file(rel_path, abs_path, progress)
        if qa_count > 0:
            total_qa += qa_count
            processed += 1
        elif progress["completed"][rel_path]["status"] in ("api_error", "error"):
            errors += 1

        progress["total_qa"] = total_qa

        # Save progress periodically (every 10 files)
        if (i + 1) % 10 == 0:
            save_progress(progress)
            print(f"  --- checkpoint: {total_qa} total Q&A, {processed} done, {errors} errors ---")

        # Small delay to avoid rate limits
        if (i + 1) % 5 == 0:
            time.sleep(0.5)

    # Final save
    save_progress(progress)

    print(f"\n{'='*60}")
    print(f"Done! Total Q&A pairs: {total_qa}")
    print(f"Files with Q&A: {processed}")
    print(f"Errors: {errors}")
    print(f"Output directory: {OUTPUT_ROOT}")
    print(f"Progress saved to: {PROGRESS_FILE}")

if __name__ == '__main__':
    main()
