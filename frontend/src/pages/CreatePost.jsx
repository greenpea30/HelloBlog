import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { createPost } from '../api/client.js'

export default function CreatePost() {
  const [title, setTitle] = useState('')
  const [summary, setSummary] = useState('')
  const [content, setContent] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  async function handleSubmit(e) {
    e.preventDefault()
    setError('')
    try {
      const data = await createPost(title, summary, content)
      navigate(`/post/${data.id}`)
    } catch (err) {
      setError(err.message)
    }
  }

  const token = localStorage.getItem('token')
  if (!token) {
    navigate('/login')
    return null
  }

  return (
    <div>
      <h1 style={{ marginBottom: 24 }}>写文章</h1>
      {error && <p style={{ color: '#e74c3c', marginBottom: 16 }}>{error}</p>}
      <form onSubmit={handleSubmit} className="post-detail">
        <div className="form-group">
          <label>标题</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            maxLength={200}
            required
            placeholder="文章标题"
          />
        </div>
        <div className="form-group">
          <label>摘要（可选）</label>
          <input
            type="text"
            value={summary}
            onChange={(e) => setSummary(e.target.value)}
            maxLength={500}
            placeholder="简短摘要"
          />
        </div>
        <div className="form-group">
          <label>正文（支持 Markdown）</label>
          <textarea
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required
            placeholder="开始写作..."
          />
        </div>
        <button className="btn btn-primary" type="submit">发布文章</button>
      </form>
    </div>
  )
}
