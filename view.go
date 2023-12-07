package gorms

import "gorm.io/gorm"

type IView interface {
	Name() string
	RawSQL() string      // required subquery.
	Replace() bool       //If true, exec `CREATE`. If false, exec `CREATE OR REPLACE`
	CheckOption() string // optional. e.g. `WITH [ CASCADED | LOCAL ] CHECK OPTION`
}

func MigrateView(db *gorm.DB, views ...IView) error {
	for _, view := range views {
		if err := db.Migrator().CreateView(view.Name(), gorm.ViewOption{
			Replace:     view.Replace(),
			CheckOption: view.CheckOption(),
			Query:       db.Raw(view.RawSQL()),
		}); err != nil {
			return err
		}
	}
	return nil
}
