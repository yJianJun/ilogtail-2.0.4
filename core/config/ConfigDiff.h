/*
 * Copyright 2023 iLogtail Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#pragma once

#include <string>
#include <vector>

#include "config/Config.h"

namespace logtail {

struct ConfigDiff {
    std::vector<Config> mAdded;
    std::vector<Config> mModified;
    std::vector<std::string> mRemoved;
    std::vector<std::string> mUnchanged; // 过渡使用，仅供插件系统用

    bool IsEmpty() { return mRemoved.empty() && mAdded.empty() && mModified.empty(); }
};

} // namespace logtail
