import { useState, useEffect } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { getPosts, search, toggleLike } from '../api/client.js'

export default function Home() {
  const [posts, setPosts] = useState([])
  const [query, setQuery] = useState('')
  const [page, setPage] = useState(1)
  const [total, setTotal] = useState(0)
  const [orderBy, setOrderBy] = useState('')
  const [zjuOnly, setZjuOnly] = useState(false)
  const [likedIds, setLikedIds] = useState(new Set())
  const navigate = useNavigate()
  const token = localStorage.getItem('token')
  const user = JSON.parse(localStorage.getItem('user') || '{}')

  const pageSize = 20

  useEffect(() => {
    loadPosts()
  }, [page, orderBy, zjuOnly])

  useEffect(() => {
    if (!token) return
    fetch('/api/v1/likes/user-liked-posts', {
      headers: { 'Authorization': `Bearer ${token}` }
    })
    .then(r => r.json())
    .then(d => {
      if (d.data?.liked_post_ids) setLikedIds(new Set(d.data.liked_post_ids))
    })
    .catch(() => {})
  }, [token])

  async function loadPosts() {
    try {
      const data = await getPosts(page, pageSize, orderBy, zjuOnly)
      setPosts(data.items)
      setTotal(data.total)
    } catch (err) {
      console.error('load posts failed:', err)
    }
  }

  async function handleSearch(e) {
    e.preventDefault()
    if (!query.trim()) {
      loadPosts()
      return
    }
    try {
      const data = await search(query.trim())
      const items = data.items || []
      setPosts(items)
      setTotal(items.length)
    } catch (err) {
      console.error('search failed:', err)
    }
  }

  const totalPages = Math.ceil(total / pageSize)

  return (
    <div>
      <form className="search-bar" onSubmit={handleSearch}>
        <input
          type="text"
          placeholder="搜索文章..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
      </form>

      <div style={{ marginBottom: 16, display: 'flex', gap: 8 }}>
        <button
          className={`btn ${orderBy === '' ? 'btn-primary' : ''}`}
          onClick={() => { setOrderBy(''); setPage(1) }}
        >最新</button>
        <button
          className={`btn ${orderBy === 'popular' ? 'btn-primary' : ''}`}
          onClick={() => { setOrderBy('popular'); setPage(1) }}
        >热门</button>
        <button
          className={`btn ${zjuOnly ? 'btn-primary' : ''}`}
          onClick={() => { setZjuOnly(!zjuOnly); setPage(1) }}
          title="筛选浙大学号登录用户的帖子"
        >🎓 校友</button>
      </div>

      {posts.map((post) => {
        const postId = post.id || post.post_id
        const isLiked = likedIds.has(postId)
        return (
          <div key={postId} className="post-card">
            <h2>
              <Link to={`/post/${postId}`}>{post.title}</Link>
            </h2>
            <div className="post-meta" style={{display:'flex',alignItems:'center',gap:8}}>
              {post.user?.avatar_url ? (
                <Link to={`/user/${post.user.id}`}><img src={post.user.avatar_url} className="avatar-sm" alt="" /></Link>
              ) : (
                <Link to={`/user/${post.user.id}`} className="avatar-sm avatar-sm-text">{post.user?.username?.[0]?.toUpperCase() || '?'}</Link>
              )}
              <Link to={`/user/${post.user.id}`} style={{ color: '#555', textDecoration: 'none' }}>{post.user?.username || '未知'}</Link> · {new Date(post.created_at).toLocaleDateString('zh-CN')} · 👁 {post.view_count ?? 0} · 💬 {post.comment_count ?? 0}
            </div>
            {post.summary && <p className="post-summary">{post.summary}</p>}
            <div className="post-actions">
              <span
                className={`like-btn ${isLiked ? 'liked' : ''}`}
                onClick={async (e) => {
                  e.preventDefault()
                  if (!token) { navigate('/login'); return }
                  try {
                    const res = await toggleLike('post', postId)
                    setLikedIds(prev => {
                      const next = new Set(prev)
                      if (res.liked) next.add(postId); else next.delete(postId)
                      return next
                    })
                    setPosts(prev => prev.map(p => {
                      const pid = p.id || p.post_id
                      if (pid === postId) {
                        return { ...p, like_count: (p.like_count || 0) + (res.liked ? 1 : -1) }
                      }
                      return p
                    }))
                  } catch (err) {
                    console.error('like failed:', err)
                  }
                }}
              >
                {isLiked ? '❤️' : '🤍'} {post.like_count ?? 0}
              </span>
            </div>
          </div>
        )
      })}

      {posts.length === 0 && <p style={{ textAlign: 'center', color: '#999' }}>还没有文章，快去写第一篇吧！</p>}

      {totalPages > 1 && (
        <div className="pagination">
          <button className="btn" disabled={page <= 1} onClick={() => setPage(page - 1)}>上一页</button>
          <span style={{ lineHeight: '36px' }}>{page} / {totalPages}</span>
          <button className="btn" disabled={page >= totalPages} onClick={() => setPage(page + 1)}>下一页</button>
        </div>
      )}
    </div>
  )
}
