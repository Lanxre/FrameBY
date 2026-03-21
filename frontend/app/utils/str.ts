import type { FramebyAppRole } from "~/types/frontend/enums/role";

export const formatDate = (date: Date | string) => {
  return new Date(date).toLocaleDateString("ru-RU", {
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
    case "user":
      return "Пользователь";
    case "university":
      return "Представитель университета";
    case "student":
      return "Студент";
    case "brsm":
      return "Представитель БРСМ";
    case "customer":
      return "Представитель компании";
    default:
      return "Неизвестная роль";
  }
};
