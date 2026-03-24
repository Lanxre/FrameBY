import { FramebyAppRole } from "@/types/frontend/enums/role";
import { BASE_URL } from "@/const";

export const formatDate = (date: Date | string) => {
  const formattedDate = new Date(date)
  formattedDate.setHours(formattedDate.getHours() - 3)   
    
  return formattedDate.toLocaleDateString("ru-RU", {
    year: "numeric",
    month: "long",
    day: "numeric",
    hour: "numeric",
    minute: "numeric",
    second: "numeric",
  });
};

export const formatRole = (role: FramebyAppRole) => {
  switch (role) {
    case FramebyAppRole.USER:
      return "Пользователь";
    case FramebyAppRole.UNIVERSITY:
      return "Представитель университета";
    case FramebyAppRole.STUDENT:
      return "Студент";
    case FramebyAppRole.BRSM:
      return "Представитель БРСМ";
    case FramebyAppRole.CUSTOMER:
      return "Представитель компании";
    case FramebyAppRole.ADMIN:
      return "Администратор";
    default:
      return "Неизвестная роль";
  }
};

export const formatAvatar = (login: string, avatar?: string) => {
  if (!avatar) return login.charAt(0).toUpperCase();
  return `${BASE_URL}/uploads/avatars/${avatar}`;
};
