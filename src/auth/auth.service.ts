import { HttpException, Injectable } from '@nestjs/common';
import { RegisterDto } from './dto/register.dto';
import { LoginDto } from './dto/login.dto';
import { UsersService } from 'src/users/users.service';
import * as bcrypt from 'bcrypt';
import { JwtService } from '@nestjs/jwt';

@Injectable()
export class AuthService {
  constructor(
    private readonly userService: UsersService,
    private readonly jwtsService: JwtService,
  ) {}
  async register(registerDto: RegisterDto) {
    const user = await this.userService.findUserByEmail(registerDto.email);
    if (user) {
      throw new HttpException('user already exists', 400);
    }
    registerDto.password = await bcrypt.hash(registerDto.password, 10);
    return await this.userService.create(registerDto);
  }
  async login(loginDto: LoginDto) {
    const user = await this.userService.findUserByEmail(loginDto.email);
    if (!user) {
      throw new HttpException('user not found', 404);
    }
    const isPasswordMatch = await bcrypt.compare(
      loginDto.password,
      user.password,
    );
    // console.log(isPasswordMatch);
    if (!isPasswordMatch) {
      throw new HttpException('wrong password', 400);
    }
    const accessToken = this.jwtsService.sign({
      sub: user.id,
      email: user.email,
    });
    return {
      accessToken: accessToken,
    };
  }
}
