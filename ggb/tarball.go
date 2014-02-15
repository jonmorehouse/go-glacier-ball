package ggb

type Tarball struct {

	id int
	key string
	Full bool
	currentSize int64 //bytes
}

func NewTarball(id int64) *Tarball {

	// createa  
	return nil
}

func (t *Tarball) AddFile(file *File) error {
		
	return nil
}

func (t *Tarball) Upload() {

	// upload the tarball

}

func (t *Tarball) Save(path string) {


}

