/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * ctrlX CORE - Data Layer API
 * This is the base API for the ctrlX Data Layer access on ctrlX CORE devices <ul> <li>Click 'Authorize' to open the 'Available authorizations' dialog.</li> <li>Enter 'username' and 'password'. The 'Client credentials location' selector together with the 'client_id' and 'client_secret' fields as well as the 'Bearer' section can be ignored.</li> <li>Click 'Authorize' and then 'Close' to close the 'Available authorizations' dialog.</li> <li>Try out those GET, PUT, ... operations you're interested in.</li> </ul>
 *
 * The version of the OpenAPI document: 2.1.0
 * Contact: support@boschrexroth.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.apache.plc4x.java.ctrlx.readwrite.rest.datalayer.model;

import com.fasterxml.jackson.annotation.*;

import java.io.UnsupportedEncodingException;
import java.net.URI;
import java.net.URLEncoder;
import java.util.*;

/**
 * This document defines a \&quot;problem detail\&quot; as a way  to carry machine-readable details of errors in a  HTTP response to avoid the need to define new error  response formats for HTTP APIs. 
 */
@JsonPropertyOrder({
  Problem.JSON_PROPERTY_TYPE,
  Problem.JSON_PROPERTY_TITLE,
  Problem.JSON_PROPERTY_STATUS,
  Problem.JSON_PROPERTY_DETAIL,
  Problem.JSON_PROPERTY_INSTANCE,
  Problem.JSON_PROPERTY_CODE,
  Problem.JSON_PROPERTY_MAIN_DIAGNOSIS_CODE,
  Problem.JSON_PROPERTY_DETAILED_DIAGNOSIS_CODE,
  Problem.JSON_PROPERTY_DYNAMIC_DESCRIPTION,
  Problem.JSON_PROPERTY_SEVERITY,
  Problem.JSON_PROPERTY_LINKS,
  Problem.JSON_PROPERTY_MORE_INFO
})
@jakarta.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-11-18T13:34:36.056861+01:00[Europe/Berlin]")
public class Problem {
  public static final String JSON_PROPERTY_TYPE = "type";
  private URI type = URI.create("about:blank");

  public static final String JSON_PROPERTY_TITLE = "title";
  private String title;

  public static final String JSON_PROPERTY_STATUS = "status";
  private Integer status;

  public static final String JSON_PROPERTY_DETAIL = "detail";
  private String detail;

  public static final String JSON_PROPERTY_INSTANCE = "instance";
  private String instance;

  public static final String JSON_PROPERTY_CODE = "code";
  private String code;

  public static final String JSON_PROPERTY_MAIN_DIAGNOSIS_CODE = "mainDiagnosisCode";
  private String mainDiagnosisCode;

  public static final String JSON_PROPERTY_DETAILED_DIAGNOSIS_CODE = "detailedDiagnosisCode";
  private String detailedDiagnosisCode;

  public static final String JSON_PROPERTY_DYNAMIC_DESCRIPTION = "dynamicDescription";
  private String dynamicDescription;

  /**
   *  Severity of a problem as defined RFC5424 of the Syslog standard, see https://tools.ietf.org/html/rfc5424
   */
  public enum SeverityEnum {
    EMERGENCY("Emergency"),
    
    ALERT("Alert"),
    
    CRITICAL("Critical"),
    
    ERROR("Error"),
    
    WARNING("Warning"),
    
    NOTICE("Notice"),
    
    INFORMATIONAL("Informational"),
    
    DEBUG("Debug");

    private String value;

    SeverityEnum(String value) {
      this.value = value;
    }

    @JsonValue
    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    @JsonCreator
    public static SeverityEnum fromValue(String value) {
      for (SeverityEnum b : SeverityEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }
  }

  public static final String JSON_PROPERTY_SEVERITY = "severity";
  private SeverityEnum severity;

  public static final String JSON_PROPERTY_LINKS = "links";
  private List<URI> links = new ArrayList<>();

  public static final String JSON_PROPERTY_MORE_INFO = "moreInfo";
  private Map<String, Object> moreInfo = new HashMap<>();

  public Problem() {
  }

  public Problem type(URI type) {
    
    this.type = type;
    return this;
  }

   /**
   * A URI reference [RFC3986] that identifies the problem type.  This specification encourages that, when dereferenced,  it provide human-readable documentation for the problem type  (e.g., using HTML [W3C.REC-html5-20141028]). When this member  is not present, its value is assumed to be \&quot;about:blank\&quot;. 
   * @return type
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_TYPE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public URI getType() {
    return type;
  }


  @JsonProperty(JSON_PROPERTY_TYPE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setType(URI type) {
    this.type = type;
  }


  public Problem title(String title) {
    
    this.title = title;
    return this;
  }

   /**
   * A short, human-readable summary of the problem type.  It SHOULD NOT change from occurrence to occurrence of the problem,  except for purposes of localization (e.g., using proactive content  negotiation; see [RFC7231], Section 3.4). 
   * @return title
  **/
  @jakarta.annotation.Nonnull
  @JsonProperty(JSON_PROPERTY_TITLE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)

  public String getTitle() {
    return title;
  }


  @JsonProperty(JSON_PROPERTY_TITLE)
  @JsonInclude(value = JsonInclude.Include.ALWAYS)
  public void setTitle(String title) {
    this.title = title;
  }


  public Problem status(Integer status) {
    
    this.status = status;
    return this;
  }

   /**
   * The HTTP status code ([RFC7231], Section 6) generated by the  origin server for this occurrence of the problem. 
   * @return status
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_STATUS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public Integer getStatus() {
    return status;
  }


  @JsonProperty(JSON_PROPERTY_STATUS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setStatus(Integer status) {
    this.status = status;
  }


  public Problem detail(String detail) {
    
    this.detail = detail;
    return this;
  }

   /**
   * A human-readable explanation specific to this occurrence of the problem. 
   * @return detail
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_DETAIL)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getDetail() {
    return detail;
  }


  @JsonProperty(JSON_PROPERTY_DETAIL)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDetail(String detail) {
    this.detail = detail;
  }


  public Problem instance(String instance) {
    
    this.instance = instance;
    return this;
  }

   /**
   * A URI reference that identifies the specific occurrence of the problem. It may or may not yield further information if dereferenced. 
   * @return instance
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_INSTANCE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getInstance() {
    return instance;
  }


  @JsonProperty(JSON_PROPERTY_INSTANCE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setInstance(String instance) {
    this.instance = instance;
  }


  public Problem code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * An error code issued by the system that caused the original problem. This code can be used to track down the root cause of the error. 
   * @return code
   * @deprecated
  **/
  @Deprecated
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_CODE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getCode() {
    return code;
  }


  @JsonProperty(JSON_PROPERTY_CODE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setCode(String code) {
    this.code = code;
  }


  public Problem mainDiagnosisCode(String mainDiagnosisCode) {
    
    this.mainDiagnosisCode = mainDiagnosisCode;
    return this;
  }

   /**
   * The main diagnosis code is issued by the system that caused the problem. This code can be used to track down the root cause and source of the error. It can be used to search in the documentation for a solution. It SHOULD NOT change from occurrence to occurrence of the same problem. 
   * @return mainDiagnosisCode
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_MAIN_DIAGNOSIS_CODE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getMainDiagnosisCode() {
    return mainDiagnosisCode;
  }


  @JsonProperty(JSON_PROPERTY_MAIN_DIAGNOSIS_CODE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setMainDiagnosisCode(String mainDiagnosisCode) {
    this.mainDiagnosisCode = mainDiagnosisCode;
  }


  public Problem detailedDiagnosisCode(String detailedDiagnosisCode) {
    
    this.detailedDiagnosisCode = detailedDiagnosisCode;
    return this;
  }

   /**
   * The detailed diagnosis code issued by the system that caused the problem. This code can be used to track down the detailed cause and source of the error. It can be used to search in the documentation for a solution. The detailed diagnosis code gives additional information about the cause of the error. It SHOULD NOT change from occurrence to occurrence of the same problem. 
   * @return detailedDiagnosisCode
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_DETAILED_DIAGNOSIS_CODE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getDetailedDiagnosisCode() {
    return detailedDiagnosisCode;
  }


  @JsonProperty(JSON_PROPERTY_DETAILED_DIAGNOSIS_CODE)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDetailedDiagnosisCode(String detailedDiagnosisCode) {
    this.detailedDiagnosisCode = detailedDiagnosisCode;
  }


  public Problem dynamicDescription(String dynamicDescription) {
    
    this.dynamicDescription = dynamicDescription;
    return this;
  }

   /**
   * A dynamic description gives detailed information about the occurrence of a problem. It can change between different occurrences of the same error. 
   * @return dynamicDescription
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_DYNAMIC_DESCRIPTION)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public String getDynamicDescription() {
    return dynamicDescription;
  }


  @JsonProperty(JSON_PROPERTY_DYNAMIC_DESCRIPTION)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setDynamicDescription(String dynamicDescription) {
    this.dynamicDescription = dynamicDescription;
  }


  public Problem severity(SeverityEnum severity) {
    
    this.severity = severity;
    return this;
  }

   /**
   *  Severity of a problem as defined RFC5424 of the Syslog standard, see https://tools.ietf.org/html/rfc5424
   * @return severity
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_SEVERITY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public SeverityEnum getSeverity() {
    return severity;
  }


  @JsonProperty(JSON_PROPERTY_SEVERITY)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setSeverity(SeverityEnum severity) {
    this.severity = severity;
  }


  public Problem links(List<URI> links) {
    
    this.links = links;
    return this;
  }

  public Problem addLinksItem(URI linksItem) {
    if (this.links == null) {
      this.links = new ArrayList<>();
    }
    this.links.add(linksItem);
    return this;
  }

   /**
   * Collection of links to fix the problem. E.g. links to online user manual, to an online community (using tags) or a developer chat (e.g. Slack). 
   * @return links
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_LINKS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)

  public List<URI> getLinks() {
    return links;
  }


  @JsonProperty(JSON_PROPERTY_LINKS)
  @JsonInclude(value = JsonInclude.Include.USE_DEFAULTS)
  public void setLinks(List<URI> links) {
    this.links = links;
  }


  public Problem moreInfo(Map<String, Object> moreInfo) {
    
    this.moreInfo = moreInfo;
    return this;
  }

  public Problem putMoreInfoItem(String key, Object moreInfoItem) {
    if (this.moreInfo == null) {
      this.moreInfo = new HashMap<>();
    }
    this.moreInfo.put(key, moreInfoItem);
    return this;
  }

   /**
   * Get moreInfo
   * @return moreInfo
  **/
  @jakarta.annotation.Nullable
  @JsonProperty(JSON_PROPERTY_MORE_INFO)
  @JsonInclude(content = JsonInclude.Include.ALWAYS, value = JsonInclude.Include.USE_DEFAULTS)

  public Map<String, Object> getMoreInfo() {
    return moreInfo;
  }


  @JsonProperty(JSON_PROPERTY_MORE_INFO)
  @JsonInclude(content = JsonInclude.Include.ALWAYS, value = JsonInclude.Include.USE_DEFAULTS)
  public void setMoreInfo(Map<String, Object> moreInfo) {
    this.moreInfo = moreInfo;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Problem problem = (Problem) o;
    return Objects.equals(this.type, problem.type) &&
        Objects.equals(this.title, problem.title) &&
        Objects.equals(this.status, problem.status) &&
        Objects.equals(this.detail, problem.detail) &&
        Objects.equals(this.instance, problem.instance) &&
        Objects.equals(this.code, problem.code) &&
        Objects.equals(this.mainDiagnosisCode, problem.mainDiagnosisCode) &&
        Objects.equals(this.detailedDiagnosisCode, problem.detailedDiagnosisCode) &&
        Objects.equals(this.dynamicDescription, problem.dynamicDescription) &&
        Objects.equals(this.severity, problem.severity) &&
        Objects.equals(this.links, problem.links) &&
        Objects.equals(this.moreInfo, problem.moreInfo);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, title, status, detail, instance, code, mainDiagnosisCode, detailedDiagnosisCode, dynamicDescription, severity, links, moreInfo);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Problem {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    title: ").append(toIndentedString(title)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    detail: ").append(toIndentedString(detail)).append("\n");
    sb.append("    instance: ").append(toIndentedString(instance)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    mainDiagnosisCode: ").append(toIndentedString(mainDiagnosisCode)).append("\n");
    sb.append("    detailedDiagnosisCode: ").append(toIndentedString(detailedDiagnosisCode)).append("\n");
    sb.append("    dynamicDescription: ").append(toIndentedString(dynamicDescription)).append("\n");
    sb.append("    severity: ").append(toIndentedString(severity)).append("\n");
    sb.append("    links: ").append(toIndentedString(links)).append("\n");
    sb.append("    moreInfo: ").append(toIndentedString(moreInfo)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

  /**
   * Convert the instance into URL query string.
   *
   * @return URL query string
   */
  public String toUrlQueryString() {
    return toUrlQueryString(null);
  }

  /**
   * Convert the instance into URL query string.
   *
   * @param prefix prefix of the query string
   * @return URL query string
   */
  public String toUrlQueryString(String prefix) {
    String suffix = "";
    String containerSuffix = "";
    String containerPrefix = "";
    if (prefix == null) {
      // style=form, explode=true, e.g. /pet?name=cat&type=manx
      prefix = "";
    } else {
      // deepObject style e.g. /pet?id[name]=cat&id[type]=manx
      prefix = prefix + "[";
      suffix = "]";
      containerSuffix = "]";
      containerPrefix = "[";
    }

    StringJoiner joiner = new StringJoiner("&");

    // add `type` to the URL query string
    if (getType() != null) {
      try {
        joiner.add(String.format("%stype%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getType()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `title` to the URL query string
    if (getTitle() != null) {
      try {
        joiner.add(String.format("%stitle%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getTitle()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `status` to the URL query string
    if (getStatus() != null) {
      try {
        joiner.add(String.format("%sstatus%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getStatus()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `detail` to the URL query string
    if (getDetail() != null) {
      try {
        joiner.add(String.format("%sdetail%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getDetail()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `instance` to the URL query string
    if (getInstance() != null) {
      try {
        joiner.add(String.format("%sinstance%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getInstance()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `code` to the URL query string
    if (getCode() != null) {
      try {
        joiner.add(String.format("%scode%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getCode()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `mainDiagnosisCode` to the URL query string
    if (getMainDiagnosisCode() != null) {
      try {
        joiner.add(String.format("%smainDiagnosisCode%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getMainDiagnosisCode()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `detailedDiagnosisCode` to the URL query string
    if (getDetailedDiagnosisCode() != null) {
      try {
        joiner.add(String.format("%sdetailedDiagnosisCode%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getDetailedDiagnosisCode()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `dynamicDescription` to the URL query string
    if (getDynamicDescription() != null) {
      try {
        joiner.add(String.format("%sdynamicDescription%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getDynamicDescription()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `severity` to the URL query string
    if (getSeverity() != null) {
      try {
        joiner.add(String.format("%sseverity%s=%s", prefix, suffix, URLEncoder.encode(String.valueOf(getSeverity()), "UTF-8").replaceAll("\\+", "%20")));
      } catch (UnsupportedEncodingException e) {
        // Should never happen, UTF-8 is always supported
        throw new RuntimeException(e);
      }
    }

    // add `links` to the URL query string
    if (getLinks() != null) {
      for (int i = 0; i < getLinks().size(); i++) {
        if (getLinks().get(i) != null) {
          try {
            joiner.add(String.format("%slinks%s%s=%s", prefix, suffix,
                "".equals(suffix) ? "" : String.format("%s%d%s", containerPrefix, i, containerSuffix),
                URLEncoder.encode(String.valueOf(getLinks().get(i)), "UTF-8").replaceAll("\\+", "%20")));
          } catch (UnsupportedEncodingException e) {
            // Should never happen, UTF-8 is always supported
            throw new RuntimeException(e);
          }
        }
      }
    }

    // add `moreInfo` to the URL query string
    if (getMoreInfo() != null) {
      for (String _key : getMoreInfo().keySet()) {
        try {
          joiner.add(String.format("%smoreInfo%s%s=%s", prefix, suffix,
              "".equals(suffix) ? "" : String.format("%s%d%s", containerPrefix, _key, containerSuffix),
              getMoreInfo().get(_key), URLEncoder.encode(String.valueOf(getMoreInfo().get(_key)), "UTF-8").replaceAll("\\+", "%20")));
        } catch (UnsupportedEncodingException e) {
          // Should never happen, UTF-8 is always supported
          throw new RuntimeException(e);
        }
      }
    }

    return joiner.toString();
  }

}

