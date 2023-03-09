package controller

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/labstack/echo/v4"
	"github.com/pawelWritesCode/df"

	"github.com/pawelWritesCode/user-crud/internal/model"
	"github.com/pawelWritesCode/user-crud/internal/repository"
)

type Error string

// User represents controller that defines method on User entity
type User struct {
	repository repository.User
}

func NewUser(userRepository repository.User) User {
	return User{repository: userRepository}
}

type UsersResponse struct {
	XMLName xml.Name `xml:"users"`
	Users   []model.User
}

// Create holds logic of new user creation
// query param format defines response body format (available values: json, xml, yaml)
func (service User) Create(context echo.Context) error {
	format := context.QueryParam("format")
	var bindUser model.User
	if err := context.Bind(&bindUser); err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	user, err := service.repository.Create(bindUser)
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if isJSON(format) {
		return context.JSON(http.StatusCreated, user)
	}

	if isYAML(format) {
		yamlResopnse, _ := yaml.Marshal(user)
		context.Response().Header().Set("Content-Type", "application/x-yaml")
		return context.String(http.StatusCreated, string(yamlResopnse))
	}

	if isXML(format) {
		return context.XMLPretty(http.StatusCreated, user, "\t")
	}

	return context.JSON(http.StatusCreated, user)
}

// GetSingle holds logic for obtaining a single user entity
// query param format defines response body format (available values: json, xml, yaml)
// query param userId should be valid user identifier
func (service User) GetSingle(context echo.Context) error {
	uidString := context.Param("userId")
	uid, err := strconv.Atoi(uidString)
	format := context.QueryParam("format")
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	user, err := service.repository.FindOne(uid)
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusNotFound, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusNotFound, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
	}

	if isJSON(format) {
		return context.JSON(http.StatusOK, user)
	}

	if isYAML(format) {
		yamlResopnse, _ := yaml.Marshal(user)
		context.Response().Header().Set("Content-Type", "application/x-yaml")
		return context.String(http.StatusOK, string(yamlResopnse))
	}

	if isXML(format) {
		return context.XMLPretty(http.StatusOK, user, "\t")
	}

	return context.JSON(http.StatusOK, user)
}

// GetMany holds logic of obtaining all the users
// query param format defines response body format (available values: json, xml, yaml)
func (service User) GetMany(context echo.Context) error {
	users := service.repository.GetAll()

	format := context.QueryParam("format")
	if isJSON(format) {
		return context.JSON(http.StatusOK, users)
	}

	if isYAML(format) {
		yamlResopnse, _ := yaml.Marshal(users)
		context.Response().Header().Set("Content-Type", "application/x-yaml")
		return context.String(http.StatusOK, string(yamlResopnse))
	}

	if isXML(format) {
		return context.XMLPretty(http.StatusOK, UsersResponse{Users: users}, "\t")
	}

	return context.JSON(http.StatusOK, users)
}

// Replace holds logic of user replacement
// query param format defines response body format (available values: json, xml, yaml)
// query param userId should be valid user identifier
func (service User) Replace(context echo.Context) error {
	format := context.QueryParam("format")
	uidString := context.Param("userId")
	uid, err := strconv.Atoi(uidString)
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	var u model.User
	if err = context.Bind(&u); err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if err = service.repository.Replace(uid, u); err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if isJSON(format) {
		return context.JSON(http.StatusOK, map[string]interface{}{})
	}

	if isYAML(format) {
		context.Response().Header().Set("Content-Type", "application/x-yaml")
		return nil
	}

	if isXML(format) {
		return context.XMLPretty(http.StatusOK, "", "\t")
	}

	return context.JSON(http.StatusOK, map[string]interface{}{})
}

// Delete holds logic of user removal
// query param format defines response body format (available values: json, xml, yaml)
// query param userId should be valid user identifier
func (service User) Delete(context echo.Context) error {
	format := context.QueryParam("format")
	uidString := context.Param("userId")
	uid, err := strconv.Atoi(uidString)
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	if err = service.repository.Delete(uid); err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return context.NoContent(http.StatusNoContent)
}

// ReceiveAvatar holds logic of persisting user avatar
//
// query param format defines response body format (available values: json, xml, yaml)
// query param userId should be valid user identifier
//
// form should have keys
//
//	name: name of file with extension
//	avatar: avatar data of multipart type
func (service User) ReceiveAvatar(context echo.Context) error {
	format := context.QueryParam("format")
	uidString := context.Param("userId")
	uid, err := strconv.Atoi(uidString)
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	name := context.FormValue("name")
	if len(name) == 0 {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": "name should not be empty string and should contain extension part"})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": "name should not be empty string and should contain extension part"})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error("name should not be empty string and should contain extension part"), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": "name should not be empty string and should contain extension part"})
	}

	avatar, err := context.FormFile("avatar")
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	src, err := avatar.Open()
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "a" + err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": "a" + err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusInternalServerError, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusInternalServerError, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "a" + err.Error()})
	}

	defer src.Close()

	fileName := uidString + "_" + name
	user, err := service.repository.FindOne(uid)
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusNotFound, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusNotFound, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
	}

	user.Avatar = fileName[len(uidString)+1:]
	err = service.repository.Replace(uid, user)
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusBadRequest, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusBadRequest, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	dst, err := os.Create(path.Join(os.TempDir(), fileName))
	if err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "b" + err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": "b" + err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusInternalServerError, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusInternalServerError, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "b" + err.Error()})
	}

	if _, err = io.Copy(dst, src); err != nil {
		if isJSON(format) {
			return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "c" + err.Error()})
		}

		if isYAML(format) {
			yamlResopnse, _ := yaml.Marshal(map[string]interface{}{"error": "c" + err.Error()})
			context.Response().Header().Set("Content-Type", "application/x-yaml")
			return context.String(http.StatusInternalServerError, string(yamlResopnse))
		}

		if isXML(format) {
			return context.XMLPretty(http.StatusInternalServerError, Error(err.Error()), "\t")
		}

		return context.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "c" + err.Error()})
	}

	return nil
}

// isJSON checks whether provided format points at json
func isJSON(format string) bool {
	return strings.ToLower(format) == string(df.JSON)
}

// isXML checks whether provided format points at xml
func isXML(format string) bool {
	return strings.ToLower(format) == string(df.XML)
}

// isYAML checks whether provided format points at yaml
func isYAML(format string) bool {
	return strings.ToLower(format) == string(df.YAML)
}
