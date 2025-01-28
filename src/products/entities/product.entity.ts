import Users from 'src/users/entities/user.entity';
import { Column, Entity, ManyToOne, PrimaryGeneratedColumn } from 'typeorm';

@Entity('products')
export default class products {
  @PrimaryGeneratedColumn()
  id: number;
  @Column({ nullable: true })
  title: string;
  @Column({ nullable: true })
  description: string;
  @Column({ nullable: true })
  price: number;
  @ManyToOne()
  user: Users;
}
