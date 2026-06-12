import { useState, useEffect } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { getUnreadCount } from '../api/client.js'

export default function Navbar() {
  const token = localStorage.getItem('token')
  const navigate = useNavigate()
  const [unread, setUnread] = useState(0)

  useEffect(() => {
    if (!token) return
    getUnreadCount().then(data => setUnread(data.unread_count)).catch(() => {})
    const timer = setInterval(() => {
      getUnreadCount().then(data => setUnread(data.unread_count)).catch(() => {})
    }, 30000)
    return () => clearInterval(timer)
  }, [token])

  function handleLogout() {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    navigate('/')
  }

  return (
    <nav className="navbar">
      <Link to="/" className="navbar-brand">HelloBlog</Link>
      <div className="navbar-links">
        {token ? (
          <>
            <Link to="/links">🔗 友链</Link>
            <Link to="/create">写文章</Link>
            <Link to="/profile" className="nav-icon-link">
              👤 <span>我的</span>
            </Link>
            <Link to="/notifications" className="nav-icon-link">
              🔔 {unread > 0 && <span className="badge">{unread}</span>}
            </Link>
            <button className="btn btn-logout" onClick={handleLogout}>退出</button>
          </>
        ) : (
          <>
            <Link to="/login">登录</Link>
            <Link to="/register" className="btn btn-primary">注册</Link>
          </>
        )}
      </div>
    </nav>
  )
}
