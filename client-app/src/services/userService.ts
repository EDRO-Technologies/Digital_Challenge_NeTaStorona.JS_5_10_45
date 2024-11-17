import type { IUser } from "@/stores/models/User";
import axios from "axios";

export const getUser = async (): Promise<IUser> => {
  try {
    const token = localStorage.getItem('token');

    const { data } = await axios.get('http://172.29.14.189:3000/auth/me', { headers: { Authorization: `Bearer ${token}` }  });

    return {
      name: data.fio,
      email: data.email,
      role: data.is_agent ? 'Организатор' : 'Пользователь',
    } as IUser;

  } catch (error) {
    console.error(error);

    return {
      name: "Гость",
      email: "1234",
      role: "user",
    } as IUser;
  }


};
