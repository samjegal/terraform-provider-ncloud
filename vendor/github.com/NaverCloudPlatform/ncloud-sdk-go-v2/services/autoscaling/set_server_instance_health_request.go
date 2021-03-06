/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-08-07T06:47:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type SetServerInstanceHealthRequest struct {

	// 헬스상태코드
HealthStatusCode *string `json:"healthStatusCode"`

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`

	// health-check grace-period 존중여부
ShouldRespectGracePeriod *bool `json:"shouldRespectGracePeriod,omitempty"`
}
