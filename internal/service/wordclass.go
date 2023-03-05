package service

import (
	"context"

	"github.com/spongeling/admin-api/internal/dao"
)

// GetWordClass fetches all word classes from the database
func (s *Service) GetWordClass(ctx context.Context) ([]*dao.Class, error) {
	return s.repo.GetAllClasses(ctx)
}

// AddWordClass adds a new word class to the database
func (s *Service) AddWordClass(ctx context.Context, m dao.Class, words []string) (uint, error) {
	var classId uint
	err := s.repo.RunTx(ctx, func(ctx context.Context) error {
		// get the ID of the word
		w, err := s.repo.GetWords(ctx, []string{m.Name})
		if err != nil {
			return err
		}

		// set the word ID for the class
		if len(w) != 0 {
			m.WordId = &w[0].ID
		}

		// add the class to the database and get the class ID
		cId, err := s.repo.AddClass(ctx, &m)
		if err != nil {
			return err
		}

		// set the class ID
		classId = cId

		// retrieve the words from the database
		ws, err := s.repo.GetWords(ctx, words)
		if err != nil {
			return err
		}

		// create a map of the words and their IDs
		var hash = make(map[string]uint64)
		for _, w := range ws {
			hash[w.Word] = w.ID
		}

		// create a list of new words that need to be added to the database
		var newWords []string
		for _, w := range words {
			if _, ok := hash[w]; !ok {
				newWords = append(newWords, w)
			}
		}

		// add the new words to the database and get their IDs
		var newWordIds []uint64
		if len(newWords) != 0 {
			newWordIds, err = s.repo.AddWords(ctx, newWords)
			if err != nil {
				return err
			}
		}

		// create a list of WordClass objects that link the words to the class
		var wc []dao.WordClass
		for _, id := range newWordIds {
			wc = append(wc, dao.WordClass{
				WordId:  id,
				ClassId: classId,
			})
		}

		for _, w := range ws {
			wc = append(wc, dao.WordClass{
				WordId:  w.ID,
				ClassId: classId,
			})
		}

		// add the WordClass objects to the database
		return s.repo.AddWordClasses(ctx, wc)
	})
	if err != nil {
		return 0, err
	}

	return classId, nil
}

func (s *Service) UpdateWordClass(ctx context.Context, wcId uint, m dao.Class, words []string) (uint, error) {
	return 0, nil
}

// DeleteWordClass deletes a word class from the database
func (s *Service) DeleteWordClass(ctx context.Context, wcId uint) (uint, error) {
	return s.repo.DeleteClass(ctx, wcId)
}
