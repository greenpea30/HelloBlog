import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { getLinks, createLink, deleteLink } from '../api/client.js'

export default function Links() {
  const [links, setLinks] = useState([])
  const [name, setName] = useState('')
  const [url, setUrl] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()
  const token = localStorage.getItem('token')

  useEffect(() => {
    loadLinks()
  }, [])

  async function loadLinks() {
    try {
      const data = await getLinks()
      setLinks(data)
    } catch (err) {
      console.error('load links failed:', err)
    }
  }

  async function handleAdd(e) {
    e.preventDefault()
    setError('')
    if (!token) { navigate('/login'); return }
    try {
      await createLink(name.trim(), url.trim())
      setName('')
      setUrl('')
      loadLinks()
    } catch (err) {
      setError(err.message)
    }
  }

  async function handleDelete(id) {
    if (!window.confirm('确定删除这个链接吗？')) return
    try {
      await deleteLink(id)
      loadLinks()
    } catch (err) {
      alert('删除失败: ' + err.message)
    }
  }

  return (
    <div>
      <h1>🔗 友情链接</h1>

      {token && (
        <form onSubmit={handleAdd} className="link-form" style={{ margin: '24px 0' }}>
          <h3>添加链接</h3>
          {error && <p style={{ color: '#e74c3c', marginBottom: 8 }}>{error}</p>}
          <div style={{ display: 'flex', gap: 8 }}>
            <input
              type="text"
              placeholder="链接名称"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
              maxLength={100}
              style={{ flex: 1, padding: '10px 12px', border: '1px solid #d9d9d9', borderRadius: 6 }}
            />
            <input
              type="url"
              placeholder="https://example.com"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              required
              maxLength={500}
              style={{ flex: 2, padding: '10px 12px', border: '1px solid #d9d9d9', borderRadius: 6 }}
            />
            <button className="btn btn-primary" type="submit">添加</button>
          </div>
        </form>
      )}

      <div className="links-grid">
        {links.map(link => (
          <div key={link.id} className="link-card">
            <a href={link.url} target="_blank" rel="noopener noreferrer">
              <div className="link-name">{link.name}</div>
              <div className="link-url">{link.url}</div>
            </a>
            {token && (
              <button
                className="link-delete-btn"
                onClick={() => handleDelete(link.id)}
                title="删除"
              >×</button>
            )}
          </div>
        ))}
      </div>

      {links.length === 0 && (
        <p style={{ textAlign: 'center', color: '#999', marginTop: 40 }}>暂无友情链接</p>
      )}
    </div>
  )
}
