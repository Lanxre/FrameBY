export interface MenuItem {
  label: string
  to?: string
  icon?: string
  action?: () => Promise<void>
}