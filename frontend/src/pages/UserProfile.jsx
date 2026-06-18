import { useState, useEffect } from 'react'
import { useParams, Link } from 'react-router-dom'
import { getUserProfile } from '../api/client.js'

export default function UserProfile() {
  const { id } = useParams()
  const [profile, setProfile] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    setLoading(true)
    getUserProfile(id)
      .then(setProfile)
      .catch(() => setProfile(null))
      .finally(() => setLoading(false))
  }, [id])

  if (loading) return <p style={{ textAlign: 'center', marginTop: 40 }}>加载中...</p>
  if (!profile) return <p style={{ textAlign: 'center', color: '#999', marginTop: 40 }}>用户不存在</p>

  const { user, folders } = profile

  return (
    <div>
      {/* 用户信息 */}
      <div className="profile-header">
        {user.avatar_url ? (
          <img src={user.avatar_url} className="profile-avatar-img" alt=""
            onError={e => { e.target.style.display = 'none'; e.target.nextSibling.style.display = 'flex' }} />
        ) : null}
        <div className="profile-avatar" style={{ display: user.avatar_url ? 'none' : 'flex' }}>
          {user.username ? user.username[0]?.toUpperCase() : '?'}
        </div>
        <div style={{ flex: 1 }}>
          <h2 style={{ margin: 0 }}>{user.username}</h2>
          {user.bio && <p style={{ color: '#666', marginTop: 4 }}>{user.bio}</p>}
        </div>
      </div>

      {/* 按文件夹展示文章 */}
      {(!folders || folders.length === 0) && (
        <p style={{ textAlign: 'center', color: '#999', marginTop: 40 }}>还没有文章</p>
      )}

      {folders?.map(folder => (
        <div key={folder.id} style={{ marginBottom: 28 }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
            <h4 style={{ margin: 0, color: '#555' }}>📁 {folder.name}</h4>
            <span style={{ fontSize: 13, color: '#999' }}>({folder.posts.length} 篇)</span>
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
