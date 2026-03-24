export interface ProfileRow {
  user_id: string;
  email: string;
  login: string;
  role: string;
  avatar: string | null;
  full_name: string;
  subrole: string | null;
  position: string | null;
  phone: string | null;
  faculty: string | null;
  specialty: string | null;
  grade: number | null;
  department: string | null;
  enterprise_id: string | null;
  updated_at: string;
}

export interface ProfilesResponse {
  profiles: ProfileRow[];
  total: number;
  limit: number;
  offset: number;
}

export interface ProfileType {
  value: string;
  label: string;
  icon: string;
  color: string;
  bgColor: string;
}

export const PROFILE_TYPES: ProfileType[] = [
  {
    value: "all",
    label: "Все",
    icon: "ph:squares-four",
    color: "text-gray-600",
    bgColor: "bg-gray-100",
  },
  {
    value: "user",
    label: "Без профиля",
    icon: "ph:user",
    color: "text-emerald-600",
    bgColor: "bg-emerald-100",
  },
  {
    value: "student",
    label: "Студент",
    icon: "ph:student",
    color: "text-blue-600",
    bgColor: "bg-blue-100",
  },
  {
    value: "brsm",
    label: "БРСМ",
    icon: "ph:users-three",
    color: "text-purple-600",
    bgColor: "bg-purple-100",
  },
  {
    value: "university",
    label: "Университет",
    icon: "ph:graduation-cap",
    color: "text-amber-600",
    bgColor: "bg-amber-100",
  },
  {
    value: "customer",
    label: "Организация",
    icon: "ph:buildings",
    color: "text-rose-600",
    bgColor: "bg-rose-100",
  },
];

export const ROLE_LABELS: Record<string, string> = {
  student: "Студент",
  brsm: "БРСМ",
  university: "Университет",
  customer: "Организация",
};

export const ROLE_ICONS: Record<string, string> = {
  student: "ph:student",
  brsm: "ph:users-three",
  university: "ph:graduation-cap",
  customer: "ph:buildings",
};

export const SUBROLE_ROLES = ["brsm", "university"];

export const getProfileTypeInfo = (
  role: string,
): { label: string; icon: string; color: string; bgColor: string } => {
  const type = PROFILE_TYPES.find((t) => t.value === role);
  if (type) {
    return {
      label: type.label,
      icon: type.icon,
      color: type.color,
      bgColor: type.bgColor,
    };
  }
  return {
    label: role,
    icon: "ph:user",
    color: "text-gray-600",
    bgColor: "bg-gray-100",
  };
};
