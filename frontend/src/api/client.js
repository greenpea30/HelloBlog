const BASE_URL = '/api/v1'

function getToken() {
  return localStorage.getItem('token')
}

async function request(url, options = {}) {
  const token = getToken()
  const headers = {
    'Content-Type': 'application/json',
    ...options.headers,
  }
  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  const res = await fetch(`${BASE_URL}${url}`, { ...options, headers })
  const data = await res.json()

  if (data.code !== 0) {
    throw new Error(data.msg || 'request failed')
  }

  return data.data
}

// Auth
export async function register(username, email, password) {
  return request('/auth/register', {
    method: 'POST',
    body: JSON.stringify({ username, email, password }),
  })
}

export async function login(email, password) {
  return request('/auth/login', {
    method: 'POST',
    body: JSON.stringify({ email, password }),
  })
}

export async function zjuLogin(studentId, password) {
  return request('/auth/zju-login', {
    method: 'POST',
    body: JSON.stringify({ student_id: studentId, password }),
  })
}

export async function getMe() {
  return request('/users/me')
}

export async function updateProfile(username, avatarUrl, bio) {
  return request('/users/me', {
    method: 'PUT',
    body: JSON.stringify({ username, avatar_url: avatarUrl, bio }),
  })
}

// Posts
export async function getPosts(page = 1, pageSize = 20, orderBy = '', zjuOnly = false) {
  const params = new URLSearchParams({ page, page_size: pageSize })
  if (orderBy) params.set('order_by', orderBy)
  if (zjuOnly) params.set('zju_only', 'true')
  return request(`/posts?${params}`)
}

export async function getPost(id) {
  return request(`/posts/${id}`)
}

export async function createPost(title, summary, content, format = 'markdown', folderId = null) {
  return request('/posts', {
    method: 'POST',
    body: JSON.stringify({ title, summary, content, format, folder_id: folderId }),
  })
}

// Folders
export async function getFolders() {
  return request('/folders')
}

export async function createFolder(name) {
  return request('/folders', {
    method: 'POST',
    body: JSON.stringify({ name }),
  })
}

export async function deleteFolder(id) {
  return request(`/folders/${id}`, { method: 'DELETE' })
}

// User profile (public)
export async function getUserProfile(userId) {
  return request(`/users/${userId}/profile`)
}

// Comments
export async function getComments(postId) {
  return request(`/posts/${postId}/comments`)
}

export async function createComment(postId, content, parentId = null) {
  return request(`/posts/${postId}/comments`, {
    method: 'POST',
    body: JSON.stringify({ content, parent_id: parentId }),
  })
}

// Likes
export async function toggleLike(targetType, targetId) {
  return request('/likes/toggle', {
    method: 'POST',
    body: JSON.stringify({ target_type: targetType, target_id: targetId }),
  })
}

// Search
export async function search(query, page = 1) {
  const params = new URLSearchParams({ q: query, page, page_size: 20 })
  return request(`/search?${params}`)
}

// Delete post
export async function deletePost(id) {
  return request(`/posts/${id}`, { method: 'DELETE' })
}

// My posts
export async function getMyPosts(page = 1, pageSize = 20) {
  const params = new URLSearchParams({ page, page_size: pageSize })
  return request(`/users/me/posts?${params}`)
}

// Notifications
export async function getNotifications() {
  return request('/notifications')
}

export async function getUnreadCount() {
  return request('/notifications/unread-count')
}

export async function markAllRead() {
  return request('/notifications/mark-all-read', { method: 'POST' })
}

// Links
export async function getLinks() {
  return request('/links')
}

export async function createLink(name, url) {
  return request('/links', {
    method: 'POST',
    body: JSON.stringify({ name, url }),
  })
}

export async function deleteLink(id) {
  return request(`/links/${id}`, { method: 'DELETE' })
}
