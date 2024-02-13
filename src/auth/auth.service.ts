import { Injectable } from '@nestjs/common';
import { RegisterDto } from './dto/register';
import { LoginDto } from './dto/login';
import { InjectRepository } from '@nestjs/typeorm';
import Users from 'src/users/entities/user.entity';
import { Repository } from 'typeorm';
import { UsersService } from 'src/users/users.service';

@Injectable()
export class AuthService {
  constructor(
    @InjectRepository(Users)
    private readonly usersRepository: Repository<Users>,
    private readonly userService: UsersService,
  ) {}
  register(registerDto: RegisterDto) {
    const user = await this.userService
  }
  login(loginDto: LoginDto) {
    return '';
  }
}
