package service

import "demo03/com/west2/entity"

func (this *ServiceStruct) CreateNote(note entity.Note) {
	this.DB.Create(&note)
}

func (this *ServiceStruct) GetNoteById(id string) entity.Note {
	note := entity.Note{}
	this.DB.Take(&note, id)
	return note
}

func (this *ServiceStruct) UpdateNoteById(note entity.Note) {
	this.DB.Updates(note)
}

func (this *ServiceStruct) DeleteNoteById(id string) {
	this.DB.Delete(&entity.Note{}, id)
}

func (this *ServiceStruct) UpdateNoteState(id string, state string) {
	note := this.GetNoteById(id)
	note.State = state
	this.DB.Updates(note)
}

func (this *ServiceStruct) UpdateAllNoteState(state string) {
	this.DB.Model(&entity.Note{}).Where("1 = 1").Update("state", state)
}

func (this *ServiceStruct) DeleteNoteByState(state string) {
	this.DB.Where("state = ?", state).Delete(&entity.Note{})
}

func (this *ServiceStruct) DeleteAllNote() {
	this.DB.Where("1 = 1").Delete(&entity.Note{})
}

func (this *ServiceStruct) GetAllNote(pageNum int, pageSize int) []entity.Note {
	var notes []entity.Note
	this.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notes)
	return notes
}

func (this *ServiceStruct) GetNoteByState(state string, pageNum int, pageSize int) []entity.Note {
	var notes []entity.Note
	this.DB.Where("state = ?", state).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notes)
	return notes
}

func (this *ServiceStruct) SearchNote(content string, pageNum int, pageSize int) []entity.Note {
	var notes []entity.Note
	this.DB.Debug().Where("title like ? ", "%"+content+"%").Or("content like ?", "%"+content+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&notes)
	return notes
}
