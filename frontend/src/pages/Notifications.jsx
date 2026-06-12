import { useState, useEffect } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { getNotifications, markAllRead } from '../api/client.js'

export default function Notifications() {
  const [notifs, setNotifs] = useState([])
  const navigate = useNavigate()

  const token = localStorage.getItem('token')

  useEffect(() => {
    if (!token) { navigate('/login'); return }
    loadNotifs()
  }, [])

  async function loadNotifs() {
    try {
      const data = await getNotifications()
      setNotifs(data)
    } catch (err) {
      console.error('load notifications failed:', err)
    }
  }

  async function handleMarkRead() {
    try {
      await markAllRead()
      setNotifs(notifs.map(n => ({ ...n, is_read: true })))
    } catch (err) {
      console.error('mark read failed:', err)
    }
  }

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 24 }}>
        <h1>消息通知</h1>
        <button className="btn btn-primary" onClick={handleMarkRead}>全部标记已读</button>
      </div>

      {notifs.length === 0 && <p style={{ textAlign: 'center', color: '#999' }}>暂无通知</p>}

      {notifs.map(n => (
        <div key={n.id} className={`notification-item ${n.is_read ? '' : 'unread'}`}>
          <div className="notification-content">
            <div className="notification-title">
              {!n.is_read && <span className="unread-dot" />}
              {n.post_id ? (
                <Link to={`/post/${n.post_id}`}>{n.title}</Link>
              ) : (
                <span>{n.title}</span>
              )}
            </div>
            <div className="notification-text">{n.content}</div>
            <div className="notification-time">{new Date(n.created_at).toLocaleString('zh-CN')}</div>
          </div>
        </div>
      ))}
    </div>
  )
}
