package services

import "net/http"

func (contact *Contact) Get() (responseStatus int, err error) {

	db := Conect()
	defer db.Close()

	db.First(contact, contact.ID)
	db.Preload("Tematica").Find(contact)
	db.Preload("Organizacion").Find(contact)

	responseStatus = http.StatusOK
	return
}

func (contact *Contact) Create() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()

	if !db.NewRecord(contact) {
		responseStatus = http.StatusBadRequest
		return
	}

	db.Create(contact)

	if db.NewRecord(contact) {
		responseStatus = http.StatusBadRequest
		return
	}

	responseStatus = http.StatusOK
	return
}

func (contact *Contact) Update() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()
	db.Save(contact)
	responseStatus = http.StatusOK
	return
}

func (contact *Contact) Delete() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()
	db.Delete(contact)
	db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(contact)

	responseStatus = http.StatusOK
	return
}

func (contacts *Contacts) Get() (responseStatus int, err error) {
	db := Conect()
	defer db.Close()

	if contacts.IDLast == 0 {
		db.Limit(contacts.Limit).Order("id DESC").Find(&contacts.List)
	} else {
		db.Where("id < ?", contacts.IDLast).Limit(contacts.Limit).Order("id DESC").Find(&contacts.List)
	}
	for i, _ := range (*contacts).List {
		c := &(*contacts).List[i]

		if c.Idtematica.Valid {
			c.Tematica.ID = c.Idtematica.Int64
			c.Tematica.Get(db)
		}
		if c.Idorganizacion.Valid {
			c.Organizacion.ID = c.Idorganizacion.Int64
			c.Organizacion.Get(db)

		}
	}

	responseStatus = http.StatusOK
	return
}

func (contacts *Contacts) Create() (responseStatus int, err error) {
	for _, contact := range contacts.List {
		responseStatus, err = contact.Create()
		if err != nil {
			panic(err)
		}
	}
	responseStatus = http.StatusOK
	return
}

func (contacts *Contacts) Update() (responseStatus int, err error) {
	for _, contact := range contacts.List {
		responseStatus, err = contact.Update()
		if err != nil {
			panic(err)
		}
	}
	responseStatus = http.StatusOK
	return
}

func (contacts *Contacts) Delete() (responseStatus int, err error) {
	for _, contact := range contacts.List {
		responseStatus, err = contact.Delete()
		if err != nil {
			panic(err)
		}
	}
	responseStatus = http.StatusOK
	return
}
