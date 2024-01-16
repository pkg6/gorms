package gorms

import "gorm.io/gorm"

type IView interface {
	Name() string
	Query(db *gorm.DB) *gorm.DB // required subquery.
	Replace() bool              //If true, exec `CREATE`. If false, exec `CREATE OR REPLACE`
	CheckOption() string        // optional. e.g. `WITH [ CASCADED | LOCAL ] CHECK OPTION`
}

func MigrateViews(db *gorm.DB, views ...IView) error {
	for _, view := range views {
		if err := db.Migrator().CreateView(view.Name(), gorm.ViewOption{
			Replace:     view.Replace(),
			CheckOption: view.CheckOption(),
			Query:       view.Query(db),
		}); err != nil {
			return err
		}
	}
	return nil
}
