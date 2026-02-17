import api from './index'
import type {
  ApiResponse,
  LoginRequest,
  RegisterRequest,
  LoginResponse,
  ProfileResponse
} from './types'

export const authApi = {
  // User registration
  register(data: RegisterRequest): Promise<ApiResponse<LoginResponse>> {
    return api.post('/register', data)
  },

  // User login
  login(data: LoginRequest): Promise<ApiResponse<LoginResponse>> {
    return api.post('/login', data)
  },

  // Get user profile
  profile(): Promise<ApiResponse<ProfileResponse>> {
    return api.get('/profile')
  }
}