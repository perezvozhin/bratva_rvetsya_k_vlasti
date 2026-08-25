package usr_service

import "strconv"

func (s *Usr_service) CreateChatINTV(id int) {
	intID := strconv.Itoa(id)
	file, err := s.MCP.WritetoFile(intID, "CHATIDinterviews.txt")
	if err != nil {
		s.logger.Error(err)
		return
	}

	if err != nil {
		s.logger.Error("cannot create interviews_collection")
		return
	}
}
