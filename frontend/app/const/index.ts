import type { MenuItem } from "@/types/frontend/header";
import type { FooterSection } from "@/types/frontend/footer";

export const PROJECT_NAME = "FrameBy";
export const BASE_URL = "http://localhost:8080";

export const GITHUB_URL = "https://github.com/Lanxre";
export const TELEGRAM_URL = "https://t.me/Lanxre";
export const TWITTER_URL = "https://x.com/Lvnxrx";

export const REQUIRE_MENU: MenuItem[] = [
	{
		label: "Сведения о дополнительной потребности в студенческих отрядах",
		to: "/normative-sovmin#item-1",
		icon: "ph:star",
	},
	{
		label: "Заявки на создание студенческих отрядов",
		to: "/normative-sovmin#item-5",
		icon: "ph:star",
	},
	{
		label: "Окрытые заявки на вступление в студенческие отряды",
		to: "/student-squads",
		icon: "ph:star",
	},
];

export const REQUIRE_SECTION_MENU: MenuItem[] = [
	{
		label:
			"Постановление Совета Министров Республики Беларусь N 572 от 31 августа 2022 г.",
		to: "/normative-sovmin#item-14",
		icon: "ph:star",
	},
	{
		label: 'Классификатор "Специальности и квалификации" (ОКРБ 011-2022)',
		to: "/normative-okrb",
		icon: "ph:star",
	},
	{
		label:
			"Рекомендации ОАЦ при Президенте Республики Беларусь по обеспечению безопасности информации",
		to: "/normative-oac",
		icon: "ph:star",
	},
	{
		label:
			"Рекомендации по прогнозированию потребности, численности и структуры подготовки специалистов",
		to: "/normative-prognosis",
		icon: "ph:star",
	},
];

export const FOOTER_SECTIONS: FooterSection[] = [
	{
		title: "Информация",
		links: [
			{ label: "Предоставление данных", to: "/features" },
			{ label: "Ресурсы", to: "/updates" },
		],
	},
	{
		title: "Компания",
		links: [{ label: "О нас", to: "/about" }],
	},
	{
		title: "Поддержка",
		links: [
			{ label: "Документация", to: "/docs" },
			{ label: "FAQ", to: "/faq" },
			{ label: "Контакты", to: "/contact" },
		],
	},
];
