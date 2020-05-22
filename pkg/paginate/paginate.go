// Copyright 2019 pangush
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package paginate

type Paginate struct {
	Total int `json:"total"`
	CurrentPage int `json:"current_page"`
	LastPage int `json:"last_page"`
	Limit int `json:"limit"`
	Offset int `json:"-"`
}

type Config struct {
	CurrentPage int
	Limit int
}

func NewPaginate(config Config) *Paginate {
	page := config.CurrentPage
	if page <= 0 {
		page = 1
	}
	limit := config.Limit
	if config.Limit <= 0 {
		limit = 20
	}
	p := &Paginate{
		Total:       0,
		CurrentPage: page,
		LastPage:    0,
		Limit:       limit,
		Offset:      0,
	}

	p.setCurrentPage(p.CurrentPage)

	return p
}

func (p *Paginate) setCurrentPage(page int) *Paginate {
	p.CurrentPage = page
	p.Offset = p.Limit * (page -1)
	return p
}

func (p *Paginate) SetTotal(total int) *Paginate {
	lastPage := total / p.Limit
	if lastPage >= 0 {
		p.LastPage = lastPage
	}
	if total % p.Limit > 0 {
		p.LastPage = p.LastPage + 1
	}
	p.Total = total
	return p
}