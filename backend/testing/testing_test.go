package testing_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-11/backend/config"
	"github.com/rg-km/final-project-engineering-11/backend/controller"
	"github.com/rg-km/final-project-engineering-11/backend/repository"
	"github.com/rg-km/final-project-engineering-11/backend/router"
	"github.com/rg-km/final-project-engineering-11/backend/service"
)

var _ = Describe("Api", func() {
	Describe("/api/auth/login", func() {
		When("Login Valid", func() {
			It("Akan mengembalikan 200", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`)
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(200))
			})
		})
		When(" Login Tidak Valid", func() {
			It("Akan Mengembalikan 400", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1244"}`)
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(400))
			})
		})
		When("Login Valid Dan Mengembalikan Data", func() {
			It("Akan Mengembalikan Token", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`)
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(w, r)
				Expect(w.Body.String()).To(ContainSubstring("token"))
			})
		})

	})

	Describe("/api/auth/register", func() {
		When("Register Valid", func() {
			It("Akan mengembalikan 200", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{
					"username" : "halomoan46",
					"name" : "halomoan",
					"password": "123",
					"address" : "medan",
					"phone" : "0823232324",
					"email" : "halomoan@gmail.com"
					}`)
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(data))
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(200))

			})
		})
		When("Register  Username Sudah Ada", func() {
			It("Akan Mengembalikan 400", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{
					"username" : "ropel1",
					"name" : "satrio32132",
					"password": "1234",
					"address" : "BogorCuy",
					"phone" : "0823232324",
					"email" : "indogaming170@gmail.com"
					}`)
				w := httptest.NewRecorder()
				r := httptest.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(data))
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(400))
			})
		})

	})

	Describe("/api/auth/logout", func() {
		It("Akan mengembalikan 200", func() {

			db := config.GetConnection1()
			defer db.Close()
			router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/api/auth/logout", nil)
			router.ServeHTTP(w, r)
			Expect(w.Result().StatusCode).To(Equal(200))
		})
	})

	Describe("/api/user/profile", func() {
		When(" Cookie Ada", func() {
			It("Akan mengembalikan 200 Dan Mengembailkan Data User", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`)
				wr := httptest.NewRecorder()
				rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(wr, rs)
				var cookie []*http.Cookie
				for _, c := range wr.Result().Cookies() {
					cookie = append(cookie, c)
				}
				r := httptest.NewRequest("GET", "/api/user/profile", nil)
				w := httptest.NewRecorder()
				for _, c := range cookie {
					r.AddCookie(c)
				}
				router.ServeHTTP(w, r)
				userdata := UserProfileTesting{}
				json.NewDecoder(w.Body).Decode(&userdata)
				Expect(w.Result().StatusCode).To(Equal(200))
				Expect(userdata.Data.Username).To(Equal("satrio44"))
			})
		})
		When(" Cookie Tidak Ada", func() {
			It("Akan mengembalikan 401", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/user/profile", nil)
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(401))

			})
		})

	})

	Describe("api/user/mentor/detail", func() {

		When(" Cookie Tidak Ada", func() {
			It("Akan mengembalikan 401", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/user/mentor/detail/4", nil)
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(401))

			})
		})

		When(" Data Mentor ada", func() {
			It("Akan Mengembailkan 200 Dan Detail Mentor", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`)
				wr := httptest.NewRecorder()
				rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(wr, rs)
				var cookie []*http.Cookie
				for _, c := range wr.Result().Cookies() {
					cookie = append(cookie, c)
				}

				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/user/mentor/detail/4", nil)
				for _, c := range cookie {
					r.AddCookie(c)
				}
				router.ServeHTTP(w, r)
				user := UserMentorDetailTesting{}
				json.NewDecoder(w.Body).Decode(&user)
				Expect(w.Result().StatusCode).To(Equal(200))
				Expect(user.Data.Name).To(Equal("kevin"))
			})
		})
		When(" Data Mentor tidak ada", func() {
			It("Akan Mengembailkan 404", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`)
				wr := httptest.NewRecorder()
				rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(wr, rs)
				var cookie []*http.Cookie
				for _, c := range wr.Result().Cookies() {
					cookie = append(cookie, c)
				}

				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/user/mentor/detail/5", nil)
				for _, c := range cookie {
					r.AddCookie(c)
				}
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(404))
			})
		})

	})

	Describe("/api/user/update/:id", func() {
		When("Cookie Tidak Ada", func() {
			It("Akan mengembalikan 401", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := `{
					"username" : "satrio44",
					"name" : "satriowibowo",
					"password" : "1234",
					"address" : "bogor",
					"phone" : "0823232323",
					"email" : "satrio232@gmail.com"
					}`
				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/api/user/update/2", bytes.NewBuffer([]byte(data)))
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(401))
			})
		})

		When("User Mengubah Param ID Dengan ID User Lain", func() {
			It("Akan Mengembalikan 401 ", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`)
				wr := httptest.NewRecorder()
				rs := httptest.NewRequest("POST", "/api/auth/login/", bytes.NewBuffer(data))
				router.ServeHTTP(wr, rs)
				var cookie []*http.Cookie
				for _, c := range wr.Result().Cookies() {
					cookie = append(cookie, c)
				}
				data2 := `{
					"username" : "satrio44",
					"name" : "satriowibowo",
					"password" : "1234",
					"address" : "bogor",
					"phone" : "0823232323",
					"email" : "satrio232@gmail.com"
					}`
				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/api/user/update/1", bytes.NewBuffer([]byte(data2)))
				for _, c := range cookie {
					r.AddCookie(c)
				}
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(401))
			})

		})

		When("User Berhasil Mengupdate Data", func() {
			It("Akan Mengembalikan 200", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
				wr := httptest.NewRecorder()
				rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(wr, rs)
				var cookie []*http.Cookie
				for _, c := range wr.Result().Cookies() {
					cookie = append(cookie, c)
				}
				data2 := `{
					"username" : "satrio444",
					"name" :  "satriowibowo",
					"password" : "1234",
					"address" : "bogor",
					"phone" : "0823232323",
					"email" : "satrio2322@gmail.com"
					}`
				w := httptest.NewRecorder()
				r := httptest.NewRequest("PUT", "/api/user/update/2", bytes.NewBuffer([]byte(data2)))
				for _, c := range cookie {
					r.AddCookie(c)
				}
				router.ServeHTTP(w, r)
				Expect(w.Result().StatusCode).To(Equal(200))
			})
		})

		Describe("/api/user/mentor/mentorlist", func() {
			It("Akan Menampilkan data Mentor Dan 200", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
				wr := httptest.NewRecorder()
				rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(wr, rs)
				var cookie []*http.Cookie
				for _, c := range wr.Result().Cookies() {
					cookie = append(cookie, c)
				}
				mentorstruct := repository.UserRepository{Db: db}
				resmentor, _ := mentorstruct.MentorList(context.Background())
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/user/mentor/mentorlist", nil)
				for _, c := range cookie {
					r.AddCookie(c)
				}
				router.ServeHTTP(w, r)
				mentorrespon := MentorListTesting{}
				json.NewDecoder(w.Body).Decode(&mentorrespon)
				Expect(w.Result().StatusCode).To(Equal(200))
				Expect(mentorrespon.Data).To(Equal(resmentor))
			})
		})

		Describe("/api/user/mentor/mentorlist?skil", func() {
			When("Skill Mentor Tersedia", func() {
				It("Akan Mengembalikan Data Mentor Dan 200", func() {
					db := config.GetConnection1()
					defer db.Close()
					router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
					data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
					wr := httptest.NewRecorder()
					rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
					router.ServeHTTP(wr, rs)
					var cookie []*http.Cookie
					for _, c := range wr.Result().Cookies() {
						cookie = append(cookie, c)
					}
					mentorstruct := repository.UserRepository{Db: db}
					resmentor, _ := mentorstruct.GetMentorByskill(context.Background(), "Backend")
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/api/user/mentor/mentorlist?skill=Backend", nil)
					for _, c := range cookie {
						r.AddCookie(c)
					}
					router.ServeHTTP(w, r)
					mentorrespon := MentorSkillTesting{}
					json.NewDecoder(w.Body).Decode(&mentorrespon)
					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(mentorrespon.Data).To(Equal(resmentor))

				})

			})

			When("Tidak Ada Skill Mentor", func() {
				It("Akan Mengembalikan Data Kosong Dan 404", func() {
					db := config.GetConnection1()
					defer db.Close()
					router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
					data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
					wr := httptest.NewRecorder()
					rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
					router.ServeHTTP(wr, rs)
					var cookie []*http.Cookie
					for _, c := range wr.Result().Cookies() {
						cookie = append(cookie, c)
					}

					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/api/user/mentor/mentorlist?skill=Network", nil)
					for _, c := range cookie {
						r.AddCookie(c)
					}
					router.ServeHTTP(w, r)
					mentorrespon := MentorSkillTesting{}
					json.NewDecoder(w.Body).Decode(&mentorrespon)
					Expect(w.Result().StatusCode).To(Equal(404))
					Expect(mentorrespon.Data).To(BeNil())

				})

			})

		})
		Describe("/api/user/booking/mentor/:id", func() {
			When("Sukses Request Mentoring", func() {
				It("Akan Menampilkan Pesan Berhasil Mengirim Request dan 200", func() {
					db := config.GetConnection1()
					defer db.Close()
					router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
					data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
					wr := httptest.NewRecorder()
					rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
					router.ServeHTTP(wr, rs)
					var cookie []*http.Cookie
					for _, c := range wr.Result().Cookies() {
						cookie = append(cookie, c)
					}
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/api/user/booking/mentor/2", nil)
					for _, c := range cookie {
						r.AddCookie(c)
					}
					router.ServeHTTP(w, r)
					res := GeneralTesting{}
					json.NewDecoder(w.Body).Decode(&res)
					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(res.Message).To(Equal("Berhasil Mengirim Request"))

				})
			})
			When("Id Mentor Tidak Ada Didalam Database", func() {
				It("Akan Menampilkan Error Mentor NotFound dan 404", func() {
					db := config.GetConnection1()
					defer db.Close()
					router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
					data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
					wr := httptest.NewRecorder()
					rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
					router.ServeHTTP(wr, rs)
					var cookie []*http.Cookie
					for _, c := range wr.Result().Cookies() {
						cookie = append(cookie, c)
					}
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/api/user/booking/mentor/9", nil) //tidak ada mentor dengan id 9
					for _, c := range cookie {
						r.AddCookie(c)
					}
					router.ServeHTTP(w, r)
					res := GeneralErrorTesting{}
					json.NewDecoder(w.Body).Decode(&res)
					Expect(w.Result().StatusCode).To(Equal(404))
					Expect(res.Message).To(Equal("Mentor not found"))

				})
			})

		})
		Describe("/api/user/booking/status", func() {
			It("Akan Menampilkan Data List Status Request Mentoring ", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
				repodata := repository.UserRepository{Db: db}
				resdata, _ := repodata.GetAllBookStatusMemberId(context.Background(), 2)
				wr := httptest.NewRecorder()
				rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
				router.ServeHTTP(wr, rs)
				var cookie []*http.Cookie
				for _, c := range wr.Result().Cookies() {
					cookie = append(cookie, c)
				}
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/user/booking/status", nil)
				for _, c := range cookie {
					r.AddCookie(c)
				}
				router.ServeHTTP(w, r)
				res := StatusBookTesting{}
				json.NewDecoder(w.Body).Decode(&res)
				Expect(w.Result().StatusCode).To(Equal(200))
				Expect(res.Data).To(Equal(resdata))

			})

		})
		Describe("Api Search Artikel User", func() {
			Describe("/api/user/artikel", func() {
				It("Menampilkan Semua Data Artikel Dan Mengembalikan 200", func() {
					db := config.GetConnection1()
					defer db.Close()
					router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
					data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
					wr := httptest.NewRecorder()
					rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
					router.ServeHTTP(wr, rs)
					repodata := repository.UserRepository{Db: db}
					resdata, _ := repodata.GetAllArtikel(context.Background())
					var cookie []*http.Cookie
					for _, c := range wr.Result().Cookies() {
						cookie = append(cookie, c)
					}
					w := httptest.NewRecorder()
					r := httptest.NewRequest("GET", "/api/user/artikel", nil)
					for _, c := range cookie {
						r.AddCookie(c)
					}
					router.ServeHTTP(w, r)
					res := ArtikelTesting{}
					json.NewDecoder(w.Body).Decode(&res)
					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(res.Data).To(Equal(resdata))
				})
			})

			Describe("/api/user/artikel/:id", func() {
				When("Terdapat Artikel Dengan Id Yang Sama Dengan Query Param", func() {
					It("Mengembalikan Artikel Dengan Id Yang Sama Dengan Query Param Dan Mengembailkan 200", func() {
						db := config.GetConnection1()
						defer db.Close()
						router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
						data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
						wr := httptest.NewRecorder()
						rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
						router.ServeHTTP(wr, rs)
						repodata := repository.UserRepository{Db: db}
						resdata, _ := repodata.GetArtikelById(context.Background(), 1)
						var cookie []*http.Cookie
						for _, c := range wr.Result().Cookies() {
							cookie = append(cookie, c)
						}
						w := httptest.NewRecorder()
						r := httptest.NewRequest("GET", "/api/user/artikel/1", nil)
						for _, c := range cookie {
							r.AddCookie(c)
						}
						router.ServeHTTP(w, r)
						res := ArtikelDetailTesting{}
						json.NewDecoder(w.Body).Decode(&res)
						Expect(w.Result().StatusCode).To(Equal(200))
						Expect(res.Data).To(Equal(resdata))

					})

				})
				When("Artikel Tidak Ditemukan", func() {
					It("Mengembalikan 404", func() {
						db := config.GetConnection1()
						defer db.Close()
						router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
						data := []byte(`{"username":"satrio44", "password":"1234"}`) //Userid 2
						wr := httptest.NewRecorder()
						rs := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(data))
						router.ServeHTTP(wr, rs)
						var cookie []*http.Cookie
						for _, c := range wr.Result().Cookies() {
							cookie = append(cookie, c)
						}
						w := httptest.NewRecorder()
						r := httptest.NewRequest("GET", "/api/user/artikel/9", nil) // tidak ada
						for _, c := range cookie {
							r.AddCookie(c)
						}
						router.ServeHTTP(w, r)
						Expect(w.Result().StatusCode).To(Equal(404))
					})

				})

			})

		})

		Describe("/api/mentor/acc/:bookid", func() {
			It("Akan Mengupdate Status Bookingan Mentor Jika Mentor ACC", func() {
				db := config.GetConnection1()
				defer db.Close()
				router := router.Newrouter(controller.NewAuthHandler(service.NewAuthService(repository.NewUserRepository(db))))
				beforedata := GetStatusBookId(db, "HICOD0023") ///Waiting
				w := httptest.NewRecorder()
				r := httptest.NewRequest("GET", "/api/mentor/acc/HICOD0023", nil)
				router.ServeHTTP(w, r)
				afterdata := GetStatusBookId(db, "HICOD0023") // Accepted
				Expect(beforedata).To(Equal("Waiting"))
				Expect(afterdata).To(Equal("Accepted"))

			})

		})

	})

	AfterEach(func() {
		db := config.GetConnection1()
		defer db.Close()
		db.Exec("UPDATE bookmentor SET status='Waiting' WHERE bookid ='HICOD0023' ")
		db.Exec("UPDATE users SET username='satrio44' WHERE id = 2")
		db.Exec("DELETE FROM users WHERE username = 'halomoan46'")
	})

})
