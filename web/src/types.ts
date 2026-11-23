// Go 側の internal/domain の型に対応するフロントエンドの型定義。

export type QuestStatus = 'not_accepted' | 'in_progress' | 'completed' | 'failed'

export interface Delivery {
  item: string
  collected: number
  required: number
}

export interface Quest {
  id: string
  name: string
  points: number
  route: string
  deadline?: string
  status: QuestStatus
  note?: string
  delivery?: Delivery
}
