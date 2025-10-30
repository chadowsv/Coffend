export interface Food {
  food_id: number
  name: string
  description: string
  price: number
  created_at: string
  updated_at: string
  menu_id: number
}

export interface Menu {
  menu_id: number
  name: string
  menu_status: boolean
  created_at: string
  updated_at: string
  foods: Food[]
}