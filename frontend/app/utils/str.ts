import { FramebyAppRole } from "@/types/frontend/enums/role";
import { BASE_URL } from "@/const";

export const formatDate = (date: Date | string) => {
	const formattedDate = new Date(date);
	formattedDate.setHours(formattedDate.getHours() - 3);

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

export interface FormatStudentSquadOptions {
	zeroText?: string;
	showNumber?: boolean;
	oneText?: string;
	twoFourText?: string;
	fivePlusText?: string;
}

export const formatTotalStudentSquads = (
	total: number,
	options?: FormatStudentSquadOptions,
) => {
	const {
		zeroText = "Нет отрядов",
		oneText = "отряд",
		twoFourText = "отряда",
		fivePlusText = "отрядов",
		showNumber = true,
	} = options || {};

	const absTotal = Math.abs(total);

	if (absTotal === 0) {
		return zeroText;
	}

	let declension: string;

	const lastDigit = absTotal % 10;
	const lastTwoDigits = absTotal % 100;

	if (lastTwoDigits >= 11 && lastTwoDigits <= 19) {
		declension = fivePlusText;
	} else if (lastDigit === 1) {
		declension = oneText;
	} else if (lastDigit >= 2 && lastDigit <= 4) {
		declension = twoFourText;
	} else {
		declension = fivePlusText;
	}

	return showNumber ? `${absTotal} ${declension}` : declension;
};
