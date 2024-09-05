package alligator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/m41denx/alligator/options"
	"time"
)

type Location struct {
	ID        int        `json:"id"`
	Short     string     `json:"short"`
	Long      string     `json:"long"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (a *Application) ListLocations(opts ...options.ListLocationsOptions) ([]*Location, error) {
	var o string
	if opts != nil && len(opts) > 0 {
		o = options.ParseRequestOptions(&opts[0])
	}
	req := a.newRequest("GET", fmt.Sprintf("/locations?%s", o), nil)
	res, err := a.Http.Do(req)
	if err != nil {
		return nil, err
	}

	buf, err := validate(res)
	if err != nil {
		return nil, err
	}

	var model struct {
		Data []struct {
			Attributes *Location `json:"attributes"`
		} `json:"data"`
	}
	if err = json.Unmarshal(buf, &model); err != nil {
		return nil, err
	}

	locs := make([]*Location, 0, len(model.Data))
	for _, l := range model.Data {
		locs = append(locs, l.Attributes)
	}

	return locs, nil
}

func (a *Application) GetLocation(id int, opts ...options.GetLocationOptions) (*Location, error) {
	var o string
	if opts != nil && len(opts) > 0 {
		o = options.ParseRequestOptions(&opts[0])
	}
	req := a.newRequest("GET", fmt.Sprintf("/locations/%d?%s", id, o), nil)
	res, err := a.Http.Do(req)
	if err != nil {
		return nil, err
	}

	buf, err := validate(res)
	if err != nil {
		return nil, err
	}

	var model struct {
		Attributes Location `json:"attributes"`
	}
	if err = json.Unmarshal(buf, &model); err != nil {
		return nil, err
	}

	return &model.Attributes, nil
}

func (a *Application) CreateLocation(short, long string) (*Location, error) {
	data, _ := json.Marshal(map[string]string{"short": short, "long": long})
	body := bytes.Buffer{}
	body.Write(data)

	req := a.newRequest("POST", "/locations", &body)
	res, err := a.Http.Do(req)
	if err != nil {
		return nil, err
	}

	buf, err := validate(res)
	if err != nil {
		return nil, err
	}

	var model struct {
		Attributes Location `json:"attributes"`
	}
	if err = json.Unmarshal(buf, &model); err != nil {
		return nil, err
	}

	return &model.Attributes, nil
}

func (a *Application) UpdateLocation(id int, short, long string) (*Location, error) {
	data, _ := json.Marshal(map[string]string{"short": short, "long": long})
	body := bytes.Buffer{}
	body.Write(data)

	req := a.newRequest("PATCH", fmt.Sprintf("/locations/%d", id), &body)
	res, err := a.Http.Do(req)
	if err != nil {
		return nil, err
	}

	buf, err := validate(res)
	if err != nil {
		return nil, err
	}

	var model struct {
		Attributes Location `json:"attributes"`
	}
	if err = json.Unmarshal(buf, &model); err != nil {
		return nil, err
	}

	return &model.Attributes, nil
}

func (a *Application) DeleteLocation(id int) error {
	req := a.newRequest("DELETE", fmt.Sprintf("/locations/%d", id), nil)
	res, err := a.Http.Do(req)
	if err != nil {
		return err
	}

	_, err = validate(res)
	return err
}
