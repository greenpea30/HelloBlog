import { useState, useEffect, useRef } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { getMe, updateProfile, getUserProfile, deleteFolder } from '../api/client.js'

export default function Profile() {
  const [profile, setProfile] = useState(null)
  const [editing, setEditing] = useState(false)
  const [username, setUsername] = useState('')
  const [avatarUrl, setAvatarUrl] = useState('')
  const [bio, setBio] = useState('')
  const [saveMsg, setSaveMsg] = useState('')
  const [uploading, setUploading] = useState(false)
  const fileRef = useRef(null)
  const navigate = useNavigate()

  const token = localStorage.getItem('token')
  const user = JSON.parse(localStorage.getItem('user') || '{}')

  useEffect(() => {
    if (!token) { navigate('/login'); return }
    loadProfile()
  }, [])

  async function loadProfile() {
    try {
      const m = await getMe()
      setUsername(m.username)
      setAvatarUrl(m.avatar_url || '')
      setBio(m.bio || '')
      const p = await getUserProfile(m.id)
      setProfile(p)
    } catch {}
  }
  async function handleSave() {
    setSaveMsg('')
    try {
      const u = await updateProfile(username, avatarUrl, bio)
      localStorage.setItem('user', JSON.stringify(u))
      setEditing(false)
      setSaveMsg('ok')
      setTimeout(() => setSaveMsg(''), 2000)
    } catch (e) { setSaveMsg(e.message) }
  }
  async function handleFileUpload(e) {
    const file = e.target.files[0]
    if (!file) return
    setUploading(true)
    try {
      const form = new FormData()
      form.append('file', file)
      const res = await fetch('/api/v1/upload/avatar', {
        method: 'POST',
        headers: { 'Authorization': `Bearer ${token}` },
        body: form,
      })
      const data = await res.json()
      if (data.code === 0) setAvatarUrl(data.data.url)
    } catch {} finally { setUploading(false) }
  }

  const allFolders = profile?.folders || []

  return (
    <div>
      {/* 个人信息头部 */}
      <div className="profile-header">
        {avatarUrl ? (
          <img src={avatarUrl} className="profile-avatar-img" alt=""
            onError={e => { e.target.style.display = 'none'; e.target.nextSibling.style.display = 'flex' }} />
        ) : null}
        <div className="profile-avatar" style={{ display: avatarUrl ? 'none' : 'flex' }}>
          {username ? username[0]?.toUpperCase() : '?'}
        </div>
        <div style={{ flex: 1 }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
            <h2 style={{ margin: 0 }}>{username}</h2>
            <button className="btn" style={{ fontSize: 12, padding: '4px 12px' }}
              onClick={() => editing ? handleSave() : (setEditing(true), setSaveMsg(''))}>
              {editing ? '💾 保存' : '✏️ 编辑'}
            </button>
          </div>
          {bio && <p style={{ color: '#666', marginTop: 4 }}>{bio}</p>}
          {saveMsg === 'ok' && <p style={{ color: '#27ae60', marginTop: 4, fontSize: 13 }}>保存成功！</p>}
          {saveMsg && saveMsg !== 'ok' && <p style={{ color: '#e74c3c', marginTop: 4, fontSize: 13 }}>{saveMsg}</p>}
        </div>
      </div>

      {/* 编辑面板 */}
      {editing && (
        <div style={{ background: '#fff', borderRadius: 8, padding: 20, marginTop: 20, boxShadow: '0 1px 3px rgba(0,0,0,0.08)' }}>
          <div className="form-group"><label>昵称</label><input type="text" value={username} onChange={e => setUsername(e.target.value)} maxLength={50} /></div>
          <div className="form-group">
            <label>头像</label>
            <div style={{ display: 'flex', gap: 8, alignItems: 'center' }}>
              <input type="url" value={avatarUrl} onChange={e => setAvatarUrl(e.target.value)} placeholder="图片URL" style={{ flex: 1 }} />
              <span style={{ color: '#999' }}>或</span>
              <input type="file" accept="image/*" ref={fileRef} onChange={handleFileUpload} style={{ display: 'none' }} />
              <button className="btn" onClick={() => fileRef.current?.click()} disabled={uploading} style={{ whiteSpace: 'nowrap' }}>
                {uploading ? '上传中...' : '📁 本地上传'}
              </button>
            </div>
          </div>
          <div className="form-group"><label>个人简介</label><input type="text" value={bio} onChange={e => setBio(e.target.value)} maxLength={200} placeholder="介绍一下自己..." /></div>
          <div style={{ display: 'flex', gap: 8 }}><button className="btn btn-primary" onClick={handleSave}>保存修改</button><button className="btn" onClick={() => setEditing(false)}>取消</button></div>
        </div>
      )}

      {/* 按文件夹展示文章 */}
      <h3 style={{ margin: '24px 0 16px' }}>我的文章</h3>
      {allFolders.length === 0 && <p style={{ textAlign: 'center', color: '#999' }}>还没有文章</p>}

      {allFolders.map(folder => (
        <div key={folder.id} style={{ marginBottom: 28 }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
            <h4 style={{ margin: 0, color: '#555' }}>
              📁 {folder.name}
            </h4>
            <span style={{ fontSize: 13, color: '#999' }}>({folder.posts.length} 篇)</span>
            {folder.id !== 0 && (
              <button className="btn" style={{ fontSize: 11, padding: '2px 8px', color: '#e74c3c', border: '1px solid #e74c3c', background: 'transparent' }}
                onClick={async () => {
                  if (!window.confirm(`确定删除文件夹「${folder.name}」吗？文章将移回"未分类"。`)) return
                  try { await deleteFolder(folder.id); loadProfile() } catch {}
                }}>删除</button>
            )}
          </div>
          {folder.posts.map(p => (
            <div key={p.id} className="post-card" style={{ marginBottom: 8 }}>
              <h2 style={{ fontSize: 16, marginBottom: 4 }}><Link to={`/post/${p.id}`}>{p.title}</Link></h2>
              <div className="post-meta" style={{ fontSize: 13 }}>
                {new Date(p.created_at).toLocaleDateString('zh-CN')} · 👁 {p.view_count ?? 0} · ❤ {p.like_count ?? 0} · 💬 {p.comment_count ?? 0}
              </div>
              {p.summary && <p className="post-summary" style={{ fontSize: 13 }}>{p.summary}</p>}
            </div>
          ))}
        </div>
      ))}
    </div>
  )
}
