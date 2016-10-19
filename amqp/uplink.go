// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package amqp

import (
	"encoding/json"
	"fmt"
	"time"
)

// PublishUplink publishes an uplink message to the MQTT broker
func (c *DefaultPublisher) PublishUplink(dataUp UplinkMessage) error {
	key := DeviceKey{dataUp.AppID, dataUp.DevID, DeviceUplink, ""}
	msg, err := json.Marshal(dataUp)
	if err != nil {
		return fmt.Errorf("Unable to marshal the message payload")
	}
	return c.publish(key.String(), msg, time.Time(dataUp.Metadata.Time))
}
