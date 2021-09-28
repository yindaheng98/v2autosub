package main

const DefaultLinkConfTmpl = `
{
  "tag": "{{name}}",
  "protocol": "vmess",
  "settings": {
    "vnext": [
      {
        "address": "{{address}}",
        "port": {{serverPort}},
        "users": [
          {
            "id": "{{uuid}}",
            "alterId": {{aid}},
            "security": "{{security}}"
          }
        ]
      }
    ]
  },
  "streamSettings": {
    "network": "{{network}}",
    "security": "{{streamSecurity}}",
    "tlsSettings": {{tls}},
    "kcpSettings": {{kcp}},
    "wsSettings": {{ws}},
    "httpSettings": {{http}},
    "quicSettings": {{quic}},
    "mux": {
        "enabled": {{mux}},
        "concurrency": {{concurrency}}
    }
  }
}`

const DefaultLinksConfTmpl = `
{
  "log": {
    "loglevel": "warning"
  },
  "inbounds": [
    {
      "port": 1080,
      "listen": "0.0.0.0",
      "protocol": "http",
      "settings": {
        "udp": true
      }
    }
  ],
  "outbounds": [
    {{outbounds}},
    {
      "tag": "direct",
      "protocol": "freedom",
      "settings": {}
    }
  ],
  "routing": {
    "domainStrategy": "IPIfNonMatch",
    "rules": [
      {
        "type": "field",
        "balancerTag": "default-balancer",
        "domain": [
          "geosite:tld-!cn",
          "geosite:geolocation-!cn"
        ]
      },
      {
        "type": "field",
        "outboundTag": "direct",
        "domain": [
          "geosite:private",
          "geosite:cn"
        ],
        "ip": [
          "geoip:cn"
        ]
      }
    ],
    "balancers": [
      {
        "tag": "default-balancer",
        "selector": [
          {{tags}}
        ],
        "strategy": {
          "type": "random"
        }
      }
    ]
  }
}`
