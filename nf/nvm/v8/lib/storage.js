// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

'use strict';

var ContractStorage = function (handler) {
    this.storage = new Storage(handler);
};

ContractStorage.prototype = {
    delete: function (key) {
        return this.storage.del(key)
    },
    get: function (key) {
        return this.storage.get(key);
    },
    set: function (key, value) {
        return this.storage.set(key, value);
    }
};

module.exports = {
    ContractStorage: ContractStorage,
    lcs: new ContractStorage(_storage_handlers.lcs),
    gcs: new ContractStorage(_storage_handlers.gcs)
};
