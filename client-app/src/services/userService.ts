import type { IUser } from "@/stores/models/User";

export const getUser = async (): Promise<IUser> => {
  return {
    name: "Тест тестович",
    role: "user",
  } as IUser;
};
