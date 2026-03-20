import type { MenuItem } from '~/types/frontend/header';
import type { FooterSection } from '~/types/frontend/footer';

export const GITHUB_URL = 'https://github.com/Lanxre';
export const TELEGRAM_URL = 'https://t.me/Lanxre';
export const TWITTER_URL = 'https://x.com/Lvnxrx';


export const REQUIRE_MENU: MenuItem[] = [
  { label: 'Сведения о дополнительной потребности в трудовых ресурсах', to: '/normative-sovmin#item-1', icon: 'mdi:medal' },
  { label: 'Заявки на подготовку рабочих, служащих и специалистов', to: '/normative-sovmin#item-15', icon: 'mdi:medal' }
]

export const REQUIRE_SECTION_MENU: MenuItem[] = [
  { label: 'Постановление Совета Министров Республики Беларусь N 572 от 31 августа 2022 г.', to: '/normative-sovmin#item-14', icon: 'mdi:medal' },
  { label: 'Классификатор "Специальности и квалификации" (ОКРБ 011-2022)', to: '/normative-okrb', icon: 'mdi:medal' },
  { label: 'Рекомендации ОАЦ при Президенте Республики Беларусь по обеспечению безопасности информации', to: '/normative-oac', icon: 'mdi:medal' },
  { label: 'Рекомендации по прогнозированию потребности, численности и структуры подготовки специалистов', to: '/normative-prognosis', icon: 'mdi:medal' },
]

export const PROFILE_MENU: MenuItem[] = [
  { label: 'Профиль', to: '/profile', icon: 'mdi:account' },
  { label: 'Настройки', to: '/settings', icon: 'mdi:cog-outline' },
  { label: 'Выход', to: '/logout', icon: 'mdi:logout' }
]

export const FOOTER_SECTIONS: FooterSection[] = [
  {
    title: 'Продукт',
    links: [
      { label: 'Функции', to: '/features' },
      { label: 'Обновления', to: '/updates' }
    ]
  },
  {
    title: 'Компания',
    links: [
      { label: 'О нас', to: '/about' },
    ]
  },
  {
    title: 'Поддержка',
    links: [
      { label: 'Документация', to: '/docs' },
      { label: 'FAQ', to: '/faq' },
      { label: 'Контакты', to: '/contact' }
    ]
  }
]