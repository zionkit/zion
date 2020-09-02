package zconf_redis_v7

import (
	"github.com/go-redis/redis/v7"
	"github.com/zionkit/zion/conf_crypto_tls"
	"github.com/zionkit/zion/conf_time"
)

type Options struct {
	Network            string                 `json:"network"`
	Address            string                 `json:"address"`
	Database           int                    `json:"database"`
	Username           string                 `json:"username"`
	Password           string                 `json:"password"`
	MaxRetries         int                    `json:"max_retries"`
	MinRetryBackoff    conf_time.Duration     `json:"min_retry_backoff"`
	MaxRetryBackoff    conf_time.Duration     `json:"max_retry_backoff"`
	DialTimeout        conf_time.Duration     `json:"dial_timeout"`
	ReadTimeout        conf_time.Duration     `json:"read_timeout"`
	WriteTimeout       conf_time.Duration     `json:"write_timeout"`
	PoolSize           int                    `json:"pool_size"`
	MinIdleConns       int                    `json:"min_idle_conns"`
	MaxConnAge         conf_time.Duration     `json:"max_conn_age"`
	PoolTimeout        conf_time.Duration     `json:"pool_timeout"`
	IdleTimeout        conf_time.Duration     `json:"idle_timeout"`
	IdleCheckFrequency conf_time.Duration     `json:"idle_check_frequency"`
	TLSConfig          conf_crypto_tls.Config `json:"tls_config"`
}

func (opts Options) Unwrap() *redis.Options {
	return &redis.Options{
		Network:            opts.Network,
		Addr:               opts.Address,
		DB:                 opts.Database,
		Username:           opts.Username,
		Password:           opts.Password,
		MaxRetries:         opts.MaxRetries,
		MinRetryBackoff:    opts.MinRetryBackoff.Unwrap(),
		MaxRetryBackoff:    opts.MaxRetryBackoff.Unwrap(),
		DialTimeout:        opts.DialTimeout.Unwrap(),
		ReadTimeout:        opts.ReadTimeout.Unwrap(),
		WriteTimeout:       opts.WriteTimeout.Unwrap(),
		PoolSize:           opts.PoolSize,
		MinIdleConns:       opts.MinIdleConns,
		MaxConnAge:         opts.MaxConnAge.Unwrap(),
		PoolTimeout:        opts.PoolTimeout.Unwrap(),
		IdleTimeout:        opts.IdleTimeout.Unwrap(),
		IdleCheckFrequency: opts.IdleCheckFrequency.Unwrap(),
		TLSConfig:          opts.TLSConfig.Unwrap(),
	}
}

type ClusterOptions struct {
	Addresses          []string               `json:"addresses"`
	MaxRedirects       int                    `json:"max_redirects"`
	ReadOnly           bool                   `json:"read_only"`
	RouteByLatency     bool                   `json:"route_by_latency"`
	RouteRandomly      bool                   `json:"route_randomly"`
	Username           string                 `json:"username"`
	Password           string                 `json:"password"`
	MaxRetries         int                    `json:"max_retries"`
	MinRetryBackoff    conf_time.Duration     `json:"min_retry_backoff"`
	MaxRetryBackoff    conf_time.Duration     `json:"max_retry_backoff"`
	DialTimeout        conf_time.Duration     `json:"dial_timeout"`
	ReadTimeout        conf_time.Duration     `json:"read_timeout"`
	WriteTimeout       conf_time.Duration     `json:"write_timeout"`
	PoolSize           int                    `json:"pool_size"`
	MinIdleConns       int                    `json:"min_idle_conns"`
	MaxConnAge         conf_time.Duration     `json:"max_conn_age"`
	PoolTimeout        conf_time.Duration     `json:"pool_timeout"`
	IdleTimeout        conf_time.Duration     `json:"idle_timeout"`
	IdleCheckFrequency conf_time.Duration     `json:"idle_check_frequency"`
	TLSConfig          conf_crypto_tls.Config `json:"tls_config"`
}

func (opts ClusterOptions) Unwrap() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:              opts.Addresses,
		MaxRedirects:       opts.MaxRedirects,
		ReadOnly:           opts.ReadOnly,
		RouteByLatency:     opts.RouteByLatency,
		RouteRandomly:      opts.RouteRandomly,
		Username:           opts.Username,
		Password:           opts.Password,
		MaxRetries:         opts.MaxRetries,
		MinRetryBackoff:    opts.MinRetryBackoff.Unwrap(),
		MaxRetryBackoff:    opts.MaxRetryBackoff.Unwrap(),
		DialTimeout:        opts.DialTimeout.Unwrap(),
		ReadTimeout:        opts.ReadTimeout.Unwrap(),
		WriteTimeout:       opts.WriteTimeout.Unwrap(),
		PoolSize:           opts.PoolSize,
		MinIdleConns:       opts.MinIdleConns,
		MaxConnAge:         opts.MaxConnAge.Unwrap(),
		PoolTimeout:        opts.PoolTimeout.Unwrap(),
		IdleTimeout:        opts.IdleTimeout.Unwrap(),
		IdleCheckFrequency: opts.IdleCheckFrequency.Unwrap(),
		TLSConfig:          opts.TLSConfig.Unwrap(),
	}
}
