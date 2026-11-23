import { useEffect, useState } from 'react'
import './App.css'
import { fetchQuests } from './api'
import type { Quest, QuestStatus } from './types'

const STATUS_LABEL: Record<QuestStatus, string> = {
  not_accepted: '未受注',
  in_progress: '進行中',
  completed: '完了',
  failed: '失敗',
}

function DeliveryProgress({ quest }: { quest: Quest }) {
  const d = quest.delivery
  if (!d) return null
  const remaining = Math.max(d.required - d.collected, 0)
  const percent = Math.min(100, Math.round((d.collected / d.required) * 100))
  return (
    <div className="delivery">
      <div className="delivery-text">
        {d.item} {d.collected}/{d.required}
        {remaining > 0 ? `（あと ${remaining}）` : '（納品可能）'}
      </div>
      <div
        className="progress"
        role="progressbar"
        aria-valuenow={percent}
        aria-valuemin={0}
        aria-valuemax={100}
      >
        <div className="progress-bar" style={{ width: `${percent}%` }} />
      </div>
    </div>
  )
}

function QuestCard({ quest }: { quest: Quest }) {
  return (
    <li className="quest-card">
      <div className="quest-head">
        <span className="quest-name">{quest.name}</span>
        <span className={`status status-${quest.status}`}>
          {STATUS_LABEL[quest.status]}
        </span>
      </div>
      <div className="quest-meta">
        <span>{quest.points} pt</span>
        <span>{quest.route}</span>
        {quest.deadline && <span>期日: {quest.deadline}</span>}
      </div>
      {quest.note && <div className="quest-note">{quest.note}</div>}
      <DeliveryProgress quest={quest} />
    </li>
  )
}

export default function App() {
  const [quests, setQuests] = useState<Quest[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const controller = new AbortController()
    fetchQuests(controller.signal)
      .then((data) => {
        setQuests(data)
        setLoading(false)
      })
      .catch((e: unknown) => {
        if (e instanceof DOMException && e.name === 'AbortError') return
        setError(e instanceof Error ? e.message : '不明なエラーが発生しました')
        setLoading(false)
      })
    return () => controller.abort()
  }, [])

  const inProgress = quests.filter((q) => q.status === 'in_progress')
  const totalPoints = inProgress.reduce((sum, q) => sum + q.points, 0)

  return (
    <div className="app">
      <header>
        <h1>fame-quest-tracker</h1>
        <p className="subtitle">
          『ファイアーエムブレム 万紫千紅』人助け・名声管理
        </p>
        {!loading && !error && (
          <p className="summary">
            進行中 {inProgress.length} 件 ・ 獲得予定 {totalPoints} pt
          </p>
        )}
      </header>
      <main>
        {loading && <p>読み込み中…</p>}
        {error && (
          <p className="error">
            {error}
            <br />
            APIサーバー（go run ./cmd/server）が起動しているか確認してください。
          </p>
        )}
        {!loading && !error && (
          <ul className="quest-list">
            {quests.map((q) => (
              <QuestCard key={q.id} quest={q} />
            ))}
          </ul>
        )}
      </main>
    </div>
  )
}
