import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { login, zjuLogin } from '../api/client.js'

export default function Login() {
  const [mode, setMode] = useState('email')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [studentId, setStudentId] = useState('')
  const [zjuPwd, setZjuPwd] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()

  async function handleEmailLogin(e) {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      const data = await login(email, password)
      localStorage.setItem('token', data.access_token)
      localStorage.setItem('user', JSON.stringify(data.user))
      navigate('/')
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  async function handleZJULogin(e) {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      const data = await zjuLogin(studentId, zjuPwd)
      localStorage.setItem('token', data.access_token)
      localStorage.setItem('user', JSON.stringify(data.user))
      navigate('/')
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="form-card">
      <h1>登录</h1>
      <div style={{ display: 'flex', marginBottom: 24 }}>
        <button onClick={() => { setMode('email'); setError('') }} style={{ flex: 1, padding: '10px', border: '1px solid #d9d9d9', background: mode === 'email' ? '#1a73e8' : '#fff', color: mode === 'email' ? '#fff' : '#555', borderRadius: '6px 0 0 6px', cursor: 'pointer', fontWeight: 500 }}>邮箱登录</button>
        <button onClick={() => { setMode('zju'); setError('') }} style={{ flex: 1, padding: '10px', border: '1px solid #d9d9d9', background: mode === 'zju' ? '#1a73e8' : '#fff', color: mode === 'zju' ? '#fff' : '#555', borderRadius: '0 6px 6px 0', cursor: 'pointer', fontWeight: 500 }}>ZJU 学号登录</button>
      </div>
      {error && <p style={{ color: '#e74c3c', marginBottom: 16 }}>{error}</p>}
      {mode === 'email' ? (
        <form onSubmit={handleEmailLogin}>
          <div className="form-group">
            <label>邮箱</label>
            <input type="email" value={email} onChange={e => setEmail(e.target.value)} required />
          </div>
          <div className="form-group">
            <label>密码</label>
            <input type="password" value={password} onChange={e => setPassword(e.target.value)} required />
          </div>
          <button className="btn btn-primary" type="submit" style={{ width: '100%' }} disabled={loading}>{loading ? '登录中...' : '登录'}</button>
        </form>
      ) : (
        <form onSubmit={handleZJULogin}>
          <div className="form-group">
            <label>学号</label>
            <input type="text" value={studentId} onChange={e => setStudentId(e.target.value)} placeholder="浙江大学学号" required />
          </div>
          <div className="form-group">
            <label>密码</label>
            <input type="password" value={zjuPwd} onChange={e => setZjuPwd(e.target.value)} placeholder="ZJU PASS" required />
          </div>
          <button className="btn btn-primary" type="submit" style={{ width: '100%' }} disabled={loading}>{loading ? '验证中...' : '学号登录'}</button>
        </form>
      )}
      <p style={{ marginTop: 16, textAlign: 'center', fontSize: 14 }}>还没有账号？<Link to="/register">立即注册</Link></p>
    </div>
  )
}
