import { Injectable } from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { UpdateUserDto } from './dto/update-user.dto';
import { InjectRepository } from '@nestjs/typeorm';
import Users from './entities/user.entity';
import { Repository } from 'typeorm';

@Injectable()
export class UsersService {
  constructor(
    @InjectRepository(Users)
    private readonly usersRepository: Repository<Users>,
  ) {}

  findUserByEmail = async (email: string) => {
    return await this.usersRepository.findOne({
      where: { email: email },
    });
  };

  create = async (data: CreateUserDto) => {
    const user = await this.usersRepository.create(data);
    this.usersRepository.save(user);
    return user;
  };

  findAll = async () => {
    return await this.usersRepository.find();
  };

  findOne(id: number) {
    return `This action returns a #${id} user`;
  }

  update(id: number, updateUserDto: UpdateUserDto) {
    return `This action updates a #${id} user`;
  }

  remove(id: number) {
    return `This action removes a #${id} user`;
  }
}
