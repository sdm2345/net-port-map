package port_map

import "testing"

func TestRedisForward_IsMatch2(t *testing.T) {

	t.Log(len("*1\r\n$2\r\n"))
}
func TestRedisForward_IsMatch(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args:args{buf:[]byte("sdsdf")},
			want:false,
		},{
			args:args{buf:[]byte("sds*1\r\n$2\r\ndf")},
			want:false,
		},
		{
			args:args{buf:[]byte("*1\r\n$2\r\n")},
			want:true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &RedisForward{}
			if got := f.IsMatch(tt.args.buf); got != tt.want {
				t.Errorf("IsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}