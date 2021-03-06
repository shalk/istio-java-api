/*
 * *
 *  * Copyright (C) 2018 Red Hat, Inc.
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *         http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 *
 */
package me.snowdrop.istio.mixer.adapter.stdio;

/**
 * Importance level for individual items output by this adapter.
 *
 * @author <a href="claprun@redhat.com">Christophe Laprun</a>
 */
public enum Level {
    INFO(0),

    WARNING(1),

    ERROR(2);

    private final int intValue;

    Level(int intValue) {
        this.intValue = intValue;
    }

    public int value() {
        return intValue;
    }
}
