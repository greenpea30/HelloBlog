import { useState, useEffect } from 'react'
import { useParams, useNavigate, Link } from 'react-router-dom'
import ReactMarkdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import { getPost, getComments, createComment, toggleLike, deletePost } from '../api/client.js'

export default function PostDetail() {
  const { id } = useParams()
  const [post, setPost] = useState(null)
  const [comments, setComments] = useState([])
  const [commentText, setCommentText] = useState('')
  const navigate = useNavigate()

  const token = localStorage.getItem('token')
  const user = JSON.parse(localStorage.getItem('user') || '{}')

  useEffect(() => {
    loadPost()
    loadComments()
  }, [id])

  async function loadPost() {
    try {
      const data = await getPost(id)
      setPost(data)
    } catch (err) {
      console.error('load post failed:', err)
    }
  }

  async function loadComments() {
    try {
      const data = await getComments(id)
      setComments(data)
    } catch (err) {
      console.error('load comments failed:', err)
    }
  }

  async function handleLike() {
    if (!token) {
      navigate('/login')
      return
    }
    try {
      const result = await toggleLike('post', Number(id))
      setPost(prev => prev ? { ...prev, like_count: prev.like_count + (result.liked ? 1 : -1) } : prev)
    } catch (err) {
      console.error('like failed:', err)
    }
  }

  async function handleComment(e) {
    e.preventDefault()
    if (!token) {
      navigate('/login')
      return
    }
    try {
      await createComment(id, commentText)
      setCommentText('')
      loadComments()
    } catch (err) {
      console.error('comment failed:', err)
    }
  }

  if (!post) return <p>加载中...</p>

  return (
    <div>
      <div className="post-detail">
        <h1>{post.title}</h1>
        <div className="post-meta" style={{display:'flex',alignItems:'center',gap:8,flexWrap:'wrap'}}>
          {post.user?.avatar_url ? (
            <Link to={`/user/${post.user.id}`}><img src={post.user.avatar_url} className="avatar-sm" alt="" /></Link>
          ) : (
            <Link to={`/user/${post.user.id}`} className="avatar-sm avatar-sm-text">{post.user?.username?.[0]?.toUpperCase() || '?'}</Link>
          )}
          <Link to={`/user/${post.user.id}`} style={{ color: '#555', textDecoration: 'none' }}>{post.user?.username}</Link> · {new Date(post.created_at).toLocaleString('zh-CN')} ·
          👁 {post.view_count} ·
          <span style={{ cursor: 'pointer' }} onClick={handleLike}>❤ {post.like_count}</span> ·
          💬 {post.comment_count}
        </div>
        {post.summary && <blockquote style={{ color: '#666', marginTop: 16 }}>{post.summary}</blockquote>}
        {token && post.user?.id === user.id && (
          <div style={{ marginTop: 16, textAlign: 'right' }}>
            <button
              className="btn btn-danger"
              onClick={async () => {
                if (!window.confirm('确定删除这篇文章吗？')) return
                try {
                  await deletePost(id)
                  navigate('/')
                } catch (err) {
                  alert('删除失败: ' + err.message)
                }
              }}
            >删除文章</button>
          </div>
        )}
        <div className={post.format === 'markdown' ? 'content markdown-body' : 'content plain-text'}>
          {post.format === 'markdown' ? (
            <ReactMarkdown remarkPlugins={[remarkGfm]}>{post.content}</ReactMarkdown>
          ) : (
            <pre style={{ whiteSpace: 'pre-wrap', fontFamily: 'inherit', margin: 0 }}>{post.content}</pre>
          )}
        </div>
      </div>

      <div className="comments-section">
        <h3>评论 ({comments.length})</h3>
        {comments.map((c) => (
          <div key={c.id} className="comment">
            <div style={{display:'flex',alignItems:'flex-start',gap:8}}>
              {c.user?.avatar_url ? (
                <img src={c.user.avatar_url} className="avatar-xs" alt="" />
              ) : (
                <span className="avatar-xs avatar-xs-text">{c.user?.username?.[0]?.toUpperCase() || '?'}</span>
              )}
              <div>
                <span className="comment-user">{c.user?.username}</span>
                <span style={{ color: '#999', fontSize: 12, marginLeft: 8 }}>
                  {new Date(c.created_at).toLocaleString('zh-CN')}
                </span>
                <div className="comment-content">{c.content}</div>
              </div>
            </div>
          </div>
        ))}

        <form onSubmit={handleComment} style={{ marginTop: 24 }}>
          <div className="form-group">
            <textarea
              rows={3}
              placeholder="写下你的评论..."
              value={commentText}
              onChange={(e) => setCommentText(e.target.value)}
            />
          </div>
          <button className="btn btn-primary" type="submit">发表评论</button>
        </form>
      </div>
    </div>
  )
}
