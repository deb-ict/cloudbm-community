package gallery

type Database interface {
	Categories() CategoryRepository
	Media() MediaRepository
}

type CategoryRepository interface {
}

type MediaRepository interface {
}
