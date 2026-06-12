import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { register } from '../api/client.js'

export default function Register() {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  async function handleSubmit(e) {
    e.preventDefault()
    setError('')
    try {
      const data = await register(username, email, password)
      localStorage.setItem('token', data.access_token)
      localStorage.setItem('user', JSON.stringify(data.user))
      navigate('/')
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <div className="form-card">
      <h1>注册</h1>
      {error && <p style={{ color: '#e74c3c', marginBottom: 16 }}>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label>用户名</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            minLength={3}
            maxLength={50}
            required
          />
        </div>
        <div className="form-group">
          <label>邮箱（可选，ZJU 用户建议直接学号登录）</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <div className="form-group">
          <label>密码</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            minLength={8}
            required
          />
        </div>
        <button className="btn btn-primary" type="submit" style={{ width: '100%' }}>注册</button>
      </form>
      <p style={{ marginTop: 16, textAlign: 'center', fontSize: 14 }}>
        已有账号？<Link to="/login">立即登录</Link>
      </p>
    </div>
  )
}
