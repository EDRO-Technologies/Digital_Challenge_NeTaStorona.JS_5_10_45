import type { IUser } from "@/stores/models/User";

export const getUser = async (): Promise<IUser> => {
  return {
    name: "Тест тестович",
    email: "1234"
    role: "user",
  } as IUser;
};
