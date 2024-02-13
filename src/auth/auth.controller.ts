import { Controller, Post, Body } from '@nestjs/common';
import { AuthService } from './auth.service';
import { RegisterDto } from './dto/register';
import { LoginDto } from './dto/login';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post()
  register(@Body() registerDTO: RegisterDto) {
    return this.authService.register(registerDTO);
  }
  @Post()
  login(@Body() loginDTO: LoginDto) {
    return this.authService.login(loginDTO);
  }
}
