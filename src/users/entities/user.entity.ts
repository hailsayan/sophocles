import { Column, Entity, PrimaryGeneratedColumn } from 'typeorm';

@Entity('users')
export default class Users {
  @PrimaryGeneratedColumn()
  id: number;
  @Column({ unique: true, nullable: false })
  email: string;
  @Column({ length: 20, nullable: true })
  firstName: string;
  @Column({ length: 20, nullable: true })
  lastName: string;
  @Column({ nullable: true })
  age: number;
  @Column({ select: false, nullable: false })
  password: string;
}
