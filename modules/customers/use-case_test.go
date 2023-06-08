package customers

import (
	"BackendCRM/entities"
	mocks "BackendCRM/mocks/modules/customers"
	"reflect"
	"testing"
)

func TestUseCase_Create(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		user *entities.Customers
	}
	req := entities.Customers{
		FirstName: "budi",
		LastName:  "hartono",
		Email:     "budihartono@gmail.com",
		Avatar:    "bigchungus.jpg",
	}
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().Save(&req).
		Return(nil).Once()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.Save",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				user: &req,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			if err := u.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_Delete(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		id string
	}
	req := entities.Customers{
		FirstName: "Michael",
		LastName:  "Lawson",
		Email:     "budihartono@gmail.com",
		Avatar:    "bigchungus.jpg",
	}
	uid := "1"
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().SoftDel(uid).
		Return(&req, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.Delete",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				id: uid,
			},
			want:    &req,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Read(t *testing.T) {
	type fields struct {
		repo Repository
	}
	req := []entities.Customers{
		{
			FirstName: "Michael",
			LastName:  "Lawson",
			Email:     "budihartono@gmail.com",
			Avatar:    "bigchungus.jpg",
		},
	}
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().FindAll().
		Return(req, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		want    []entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.Read",
			fields: fields{
				repo: mockRepo,
			},
			want:    req,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_ReadBy(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		col string
		val string
	}
	resp := entities.Customers{
		FirstName: "Michael",
		LastName:  "Lawson",
		Email:     "budihartono@gmail.com",
		Avatar:    "bigchungus.jpg",
	}
	column := "first_name"
	value := "Michael"
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().FindBy(column, value).
		Return(&resp, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.ReadBy",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				col: column,
				val: value,
			},
			want:    &resp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.ReadBy(tt.args.col, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Update(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		cst *entities.Customers
		id  string
	}
	req := entities.Customers{
		FirstName: "Michael",
		LastName:  "Lawson",
		Email:     "budihartono@gmail.com",
		Avatar:    "bigchungus.jpg",
	}
	uid := "1"
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().ChangeById(&req, uid).
		Return(&req, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Customers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.Update",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				cst: &req,
				id:  uid,
			},
			want:    &req,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.Update(tt.args.cst, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
