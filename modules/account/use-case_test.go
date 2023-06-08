package account

import (
	"BackendCRM/entities"
	mocks "BackendCRM/mocks/modules/account"
	"reflect"
	"testing"
)

func TestUseCase_ActorReadBy(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		col string
		val string
	}
	req := entities.Actors{
		Username: "adminemo",
		Password: "thisishashed",
		RoleId:   "2",
		Verified: "true",
		Active:   "true",
	}
	column := "username"
	value := "adminemo"
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().ActorFindBy(column, value).
		Return(&req, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actors
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.FindBy",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				col: column,
				val: value,
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
			got, err := u.ActorReadBy(tt.args.col, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ActorReadBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ActorReadBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Create(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		user *entities.Actors
	}
	req := entities.Actors{
		Username: "adminemo",
		Password: "thisishashed",
		RoleId:   "2",
		Verified: "true",
		Active:   "true",
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

func TestUseCase_CreateReg(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		user *entities.Registers
	}
	req := entities.Registers{
		ID:           1,
		AdminId:      2,
		SuperAdminId: 1,
		Status:       "Active",
	}
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().SaveReg(&req).
		Return(nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.SaveReg",
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
			if err := u.CreateReg(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateReg() error = %v, wantErr %v", err, tt.wantErr)
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
	req := entities.Actors{
		Username: "adminemo",
		Password: "thisishashed",
		RoleId:   "2",
		Verified: "true",
		Active:   "true",
	}
	uid := "2"
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().SoftDel(uid).
		Return(&req, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actors
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.SoftDel",
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
	resp := []entities.Actors{
		{
			Username: "adminemo",
			Password: "thisishashed",
			RoleId:   "2",
			Verified: "true",
			Active:   "true",
		},
	}
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().FindAll().
		Return(resp, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		want    []entities.Actors
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.FindAll",
			fields: fields{
				repo: mockRepo,
			},
			want:    resp,
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

func TestUseCase_ReadRegis(t *testing.T) {
	type fields struct {
		repo Repository
	}
	resp := []entities.Registers{
		{
			ID:           1,
			AdminId:      2,
			SuperAdminId: 1,
			Status:       "Active",
		},
	}
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().FindAllRegis().
		Return(resp, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		want    []entities.Registers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.FindAllRegis",
			fields: fields{
				repo: mockRepo,
			},
			want:    resp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.ReadRegis()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadRegis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadRegis() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Update(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		cst *entities.Actors
		id  string
	}
	req := entities.Actors{
		Username: "adminemo",
		Password: "thisishashed",
		RoleId:   "2",
		Verified: "true",
		Active:   "true",
	}
	uid := "1"
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().ChangeActorById(&req, uid).
		Return(&req, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Actors
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error on repository.ChangeActorById",
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

func TestUseCase_UpdateReg(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		cst *entities.Registers
		id  string
	}
	req := entities.Registers{
		ID:           1,
		AdminId:      2,
		SuperAdminId: 1,
		Status:       "Active",
	}
	uid := "2"
	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().ChangeRegisById(&req, uid).
		Return(&req, nil).Once()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Registers
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "error in repository.ChangeRegisById",
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
			got, err := u.UpdateReg(tt.args.cst, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateReg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateReg() got = %v, want %v", got, tt.want)
			}
		})
	}
}
