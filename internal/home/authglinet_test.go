package home

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/AdguardTeam/AdGuardHome/internal/aghos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthGL(t *testing.T) {
	dir := t.TempDir()

	GLMode = true
	t.Cleanup(func() {
		GLMode = false
	})
	glFilePrefix = dir + "/gl_token_"

	data := make([]byte, 4)
	aghos.NativeEndian.PutUint32(data, 1)

	require.Nil(t, ioutil.WriteFile(glFilePrefix+"test", data, 0o644))
	assert.False(t, glCheckToken("test"))

	data = make([]byte, 4)
	aghos.NativeEndian.PutUint32(data, uint32(time.Now().UTC().Unix()+60))

	require.Nil(t, ioutil.WriteFile(glFilePrefix+"test", data, 0o644))
	r, _ := http.NewRequest(http.MethodGet, "http://localhost/", nil)
	r.AddCookie(&http.Cookie{Name: glCookieName, Value: "test"})
	assert.True(t, glProcessCookie(r))
}
