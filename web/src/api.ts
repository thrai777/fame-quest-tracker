import type { Quest } from './types'

export async function fetchQuests(signal?: AbortSignal): Promise<Quest[]> {
  const res = await fetch('/api/quests', { signal })
  if (!res.ok) {
    throw new Error(`クエストの取得に失敗しました (HTTP ${res.status})`)
  }
  return (await res.json()) as Quest[]
}
