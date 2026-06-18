import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { createPost, createFolder, getFolders } from '../api/client.js'

export default function CreatePost() {
  const [title, setTitle] = useState('')
  const [summary, setSummary] = useState('')
  const [content, setContent] = useState('')
  const [format, setFormat] = useState('markdown')
  const [folders, setFolders] = useState([])
  const [folderId, setFolderId] = useState(null)
  const [showNewFolder, setShowNewFolder] = useState(false)
  const [newFolderName, setNewFolderName] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const token = localStorage.getItem('token')
  if (!token) { navigate('/login'); return null }

  useEffect(() => {
    getFolders().then(setFolders).catch(() => {})
  }, [])

  async function handleAddFolder() {
    if (!newFolderName.trim()) return
    try {
      const f = await createFolder(newFolderName.trim())
      setFolders([...folders, f])
      setFolderId(f.id)
      setShowNewFolder(false)
      setNewFolderName('')
    } catch (err) {
      setError(err.message)
    }
  }

  async function handleSubmit(e) {
    e.preventDefault()
    setError('')
    try {
      const data = await createPost(title, summary, content, format, folderId)
      navigate(`/post/${data.id}`)
    } catch (err) {
      setError(err.message)
    }
  }

  return (
    <div>
      <h1 style={{ marginBottom: 24 }}>写文章</h1>
      {error && <p style={{ color: '#e74c3c', marginBottom: 16 }}>{error}</p>}
      <form onSubmit={handleSubmit} className="post-detail">
        <div className="form-group">
          <label>标题</label>
          <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} maxLength={200} required placeholder="文章标题" />
        </div>
        <div className="form-group">
          <label>摘要（可选）</label>
          <input type="text" value={summary} onChange={(e) => setSummary(e.target.value)} maxLength={500} placeholder="简短摘要" />
        </div>
        <div className="form-group">
          <label>正文格式</label>
          <div style={{ display: 'flex', gap: 12, marginBottom: 12 }}>
            <label style={{ cursor: 'pointer', display: 'flex', alignItems: 'center', gap: 4 }}>
              <input type="radio" name="format" value="markdown" checked={format === 'markdown'} onChange={() => setFormat('markdown')} />
              📝 Markdown
            </label>
            <label style={{ cursor: 'pointer', display: 'flex', alignItems: 'center', gap: 4 }}>
              <input type="radio" name="format" value="plain" checked={format === 'plain'} onChange={() => setFormat('plain')} />
              📄 纯文本
            </label>
          </div>
        </div>
        <div className="form-group">
          <label>文件夹</label>
          <div style={{ display: 'flex', gap: 8, alignItems: 'center' }}>
            <select
              value={folderId ?? ''}
              onChange={(e) => setFolderId(e.target.value ? Number(e.target.value) : null)}
              style={{ flex: 1, padding: '10px 12px', border: '1px solid #d9d9d9', borderRadius: 6, fontSize: 14 }}
            >
              <option value="">未分类</option>
              {folders.map(f => (
                <option key={f.id} value={f.id}>{f.name} ({f.post_count})</option>
              ))}
            </select>
            <button type="button" className="btn" onClick={() => setShowNewFolder(!showNewFolder)} style={{ whiteSpace: 'nowrap' }}>
              + 新建
            </button>
          </div>
          {showNewFolder && (
            <div style={{ display: 'flex', gap: 8, marginTop: 8 }}>
              <input
                type="text" value={newFolderName} onChange={(e) => setNewFolderName(e.target.value)}
                placeholder="文件夹名称" maxLength={50}
                style={{ flex: 1, padding: '8px 12px', border: '1px solid #d9d9d9', borderRadius: 6 }}
                onKeyDown={(e) => { if (e.key === 'Enter') { e.preventDefault(); handleAddFolder() } }}
              />
              <button type="button" className="btn btn-primary" onClick={handleAddFolder}>创建</button>
            </div>
          )}
        </div>
        <div className="form-group">
          <label>正文</label>
          <textarea value={content} onChange={(e) => setContent(e.target.value)} required
            placeholder={format === 'markdown' ? '支持 Markdown 语法...' : '请输入纯文本内容...'} />
        </div>
        <button className="btn btn-primary" type="submit">发布文章</button>
      </form>
    </div>
  )
}
