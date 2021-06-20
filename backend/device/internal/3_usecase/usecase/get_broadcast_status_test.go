package usecase

import "testing"

func TestUseCase_GetBroadcastStatus(t *testing.T) {
	type fields struct {
		ToDomain  ToDomain
		ToService ToService
	}
	type args struct {
		object string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCase{
				ToDomain:  tt.fields.ToDomain,
				ToService: tt.fields.ToService,
			}
			if got := uc.GetBroadcastStatus(tt.args.object); got != tt.want {
				t.Errorf("UseCase.GetBroadcastStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
