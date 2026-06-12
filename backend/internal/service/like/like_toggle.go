package like

func (s *Service) Toggle(userID int64, targetType string, targetID int64) (liked bool, err error) {
	liked, err = s.likes.Toggle(userID, targetType, targetID)
	if err != nil {
		return false, err
	}

	// 更新对应对象的点赞计数
	var counter LikeCounter
	switch targetType {
	case "post":
		counter = s.postCounter
	case "comment":
		counter = s.commentCounter
	}

	if counter != nil {
		if liked {
			_ = counter.IncrementLikeCount(targetID)
		} else {
			_ = counter.DecrementLikeCount(targetID)
		}
	}

	return liked, nil
}
